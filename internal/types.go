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
	Date             string
	GitlabUsername   string
	LastXDays        int
	IssueCount       int
	Issues           []Issue
	ApprovedMRsCount int
	ApprovedMRs      []MergeRequest
	MergedMRsCount   int
	MergedMRs        []MergeRequest
}

type MergeRequest struct {
	Title       string
	Description string
	Status      string
	URL         string
	MergedBy    string
	MergedAt    string
	UpdatedAt   string
}
