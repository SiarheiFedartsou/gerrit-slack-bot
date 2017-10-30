package main

import "github.com/nlopes/slack"

type Answerer interface {
	Answer(answer string)
}

type SlackChannelAnswerer struct {
	RTM     *slack.RTM
	Channel string
}

func (answerer SlackChannelAnswerer) Answer(answer string) {
	answerer.RTM.SendMessage(answerer.RTM.NewOutgoingMessage(answer, answerer.Channel))
}
