package main

import (
	"fmt"
	"strings"

	gerrit "github.com/andygrunwald/go-gerrit"
)

// AddReviewersChangeProcessor adds appropriate reviewers defined in `Config` to change
type AddReviewersChangeProcessor struct {
	Config   *Config
	Answerer Answerer
}

func (processor *AddReviewersChangeProcessor) Process(changeID string) {
	instance := processor.Config.GerritURL
	client, _ := gerrit.NewClient(instance, nil)
	client.Authentication.SetDigestAuth(processor.Config.GerritUserName, processor.Config.GerritUserPassword)
	info, _, _ := client.Changes.GetChange(changeID, nil)

	if info == nil {
		return
	}

	status := info.Status
	if status != "NEW" {
		answer := fmt.Sprintf("Invalid status of pull request #%s", changeID)
		processor.Answerer.Answer(answer)
		return
	}

	project := processor.Config.Project(info.Project)

	if project == nil {
		answer := fmt.Sprintf("No reviewers defined for %s project", info.Project)
		processor.Answerer.Answer(answer)
		return
	}

	reviewers := project.Reviewers

	if reviewers == nil {
		answer := fmt.Sprintf("No reviewers defined for %s project", info.Project)
		processor.Answerer.Answer(answer)
		return
	}

	added := []string{}
	for _, reviewer := range reviewers {
		reviewerInput := &gerrit.ReviewerInput{Reviewer: reviewer.GerritEmail, Confirmed: false}
		_, _, err := client.Changes.AddReviewer(changeID, reviewerInput)
		if err == nil {
			added = append(added, reviewer.SlackName)
		}
	}

	var answer string
	if len(added) > 0 {
		answer = "Assigned reviewers for change #" + changeID + ": " + strings.Join(added, ", ")
	}
	processor.Answerer.Answer(answer)
}
