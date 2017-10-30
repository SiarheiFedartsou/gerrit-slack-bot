package main

import (
	"encoding/json"
	"io/ioutil"
)

// Reviewer describes reviewer
type Reviewer struct {
	GerritEmail string `json:"gerrit_email"`
	SlackName   string `json:"slack_name"`
}

// Project describes project
type Project struct {
	Name      string     `json:"name"`
	Reviewers []Reviewer `json:"reviewers"`
}

// Config contains main options of the bot
type Config struct {
	Token              string    `json:"token"`
	GerritURL          string    `json:"gerrit_url"`
	GerritUserName     string    `json:"gerrit_user_name"`
	GerritUserPassword string    `json:"gerrit_user_password"`
	Projects           []Project `json:"projects"`
}

// Project finds project in config by name
func (config *Config) Project(name string) *Project {
	for _, project := range config.Projects {
		if project.Name == name {
			return &project
		}
	}
	return nil
}

// LoadConfig loads Config from file at `path`
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(data, &config)
	return config, err
}
