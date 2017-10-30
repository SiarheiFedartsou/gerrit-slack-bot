package main

import (
	"os"
	"regexp"

	"github.com/nlopes/slack"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findChangeIDs(url string, message string) []string {
	escapedURL := regexp.QuoteMeta(url)
	reString := escapedURL + ".*\\/([\\d]+)"
	re := regexp.MustCompile(reString)

	matches := re.FindAllStringSubmatch(message, -1)

	changeIds := make([]string, len(matches), len(matches))
	for index, match := range matches {
		changeIds[index] = match[1]
	}
	return changeIds
}

func handleMessage(ev *slack.MessageEvent, config *Config, rtm *slack.RTM) {
	changeIDs := findChangeIDs(config.GerritURL, ev.Text)
	for _, changeID := range changeIDs {
		answerer := SlackChannelAnswerer{RTM: rtm, Channel: ev.Channel}

		answerer.Answer("Processing #" + changeID)

		processor := NewChangeProcessor(config, answerer)
		go processor.Process(changeID)
	}
}

func main() {
	if len(os.Args) <= 1 {
		panic("No config")
	}

	config, err := LoadConfig(os.Args[1])
	check(err)

	api := slack.New(config.Token)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(ev, config, rtm)
		default:
		}
	}
}