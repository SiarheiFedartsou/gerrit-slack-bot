package main

import "regexp"

func FindChangeIDs(url string, message string) []string {
	escapedURL := regexp.QuoteMeta(url)
	reString := escapedURL + ".*?\\/([\\d]+)"
	re := regexp.MustCompile(reString)

	matches := re.FindAllStringSubmatch(message, -1)

	changeIds := make([]string, len(matches), len(matches))
	for index, match := range matches {
		changeIds[index] = match[1]
	}
	return changeIds
}
