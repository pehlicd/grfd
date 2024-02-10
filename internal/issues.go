package internal

import (
	"github.com/xanzy/go-gitlab"
	"time"
)

func (client *Client) getLastXDaysIssues(cfg Config) (*[]Issue, error) {
	updatedAfter := time.Now().Add(-24 * time.Duration(cfg.LastDays) * time.Hour)

	issues, _, err := client.Issues.ListIssues(&gitlab.ListIssuesOptions{
		AssigneeUsername: &cfg.GitlabUsername,
		UpdatedAfter:     &updatedAfter,
		OrderBy:          gitlab.Ptr("updated_at"),
		Sort:             gitlab.Ptr("desc"),
	})
	if err != nil {
		return nil, err
	}

	var is []Issue
	for _, issue := range issues {
		if issue.Description == "" {
			issue.Description = "No description provided"
		}
		i := Issue{
			Title:         issue.Title,
			URL:           issue.WebURL,
			Description:   issue.Description,
			LastUpdatedAt: issue.UpdatedAt.Format("02.01.2006 15:04:05"),
			Status:        issue.State,
		}
		is = append(is, i)
	}

	return &is, nil
}
