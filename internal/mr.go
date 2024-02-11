package internal

import (
	"github.com/xanzy/go-gitlab"
	"time"
)

func (client *Client) approveMergeRequests(cfg Config) (*[]MergeRequest, error) {
	updatedAfter := time.Now().Add(-24 * time.Duration(cfg.LastDays) * time.Hour)
	mrs, _, err := client.MergeRequests.ListMergeRequests(&gitlab.ListMergeRequestsOptions{
		ApprovedByIDs: gitlab.ApproverIDs([]int{cfg.GitlabUserId}),
		Scope:         gitlab.Ptr("all"),
		Sort:          gitlab.Ptr("desc"),
		UpdatedAfter:  &updatedAfter,
	})
	if err != nil {
		return nil, err
	}

	var approvedMRs []MergeRequest

	for _, mr := range mrs {
		if mr.Description == "" {
			mr.Description = "No description provided"
		}
		approvedMR := MergeRequest{
			Title:       mr.Title,
			Description: mr.Description,
			Status:      mr.State,
			URL:         mr.WebURL,
			UpdatedAt:   mr.UpdatedAt.Format("02.01.2006 15:04:05"),
		}
		approvedMRs = append(approvedMRs, approvedMR)
	}

	return &approvedMRs, nil
}

func (client *Client) mergedMergeRequests(cfg Config) (*[]MergeRequest, error) {
	updatedAfter := time.Now().Add(-24 * time.Duration(cfg.LastDays) * time.Hour)
	mrs, _, err := client.MergeRequests.ListMergeRequests(&gitlab.ListMergeRequestsOptions{
		UpdatedAfter: &updatedAfter,
		State:        gitlab.Ptr("merged"),
		Scope:        gitlab.Ptr("all"),
		Sort:         gitlab.Ptr("desc"),
	})
	if err != nil {
		return nil, err
	}

	var mergedMRs []MergeRequest

	for _, mr := range mrs {
		if mr.MergedBy.ID != cfg.GitlabUserId {
			continue
		}
		if mr.Description == "" {
			mr.Description = "No description provided"
		}
		mergedMR := MergeRequest{
			Title:       mr.Title,
			Description: mr.Description,
			Status:      mr.State,
			URL:         mr.WebURL,
			MergedBy:    mr.MergedBy.Name,
			MergedAt:    mr.MergedAt.Format("02.01.2006 15:04:05"),
			UpdatedAt:   mr.UpdatedAt.Format("02.01.2006 15:04:05"),
		}
		mergedMRs = append(mergedMRs, mergedMR)
	}

	return &mergedMRs, nil
}
