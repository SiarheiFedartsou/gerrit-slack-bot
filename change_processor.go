package main

// ChangeProcessor processes gerrit changes by id
type ChangeProcessor interface {
	Process(changeID string)
}

func NewChangeProcessor(config *Config, answerer Answerer) ChangeProcessor {
	return &AddReviewersChangeProcessor{Config: config, Answerer: answerer}
}
