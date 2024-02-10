package internal

import "github.com/xanzy/go-gitlab"

type Client struct {
	*gitlab.Client
}

type Config struct {
	GitLabAddr     string
	GitLabToken    string
	GitlabUsername string
	GitlabUserId   int
	LastDays       int
	TemplateFile   string
	OutputDir      string
	Insecure       bool
	Verbose        bool
}

type Issue struct {
	Title         string
	URL           string
	Description   string
	LastUpdatedAt string
	Status        string
}

type Data struct {
	Date       string
	LastXDays  int
	IssueCount int
	Issues     []Issue
}
