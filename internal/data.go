package internal

import "time"

func (client *Client) DataGenerator(cfg Config) (*Data, error) {
	date := time.Now().Format("02.01.2006")
	issues, err := client.getLastXDaysIssues(cfg)
	if err != nil {
		return nil, err
	}

	approvedMRs, err := client.approveMergeRequests(cfg)
	if err != nil {
		return nil, err
	}

	mergedMRs, err := client.mergedMergeRequests(cfg)
	if err != nil {
		return nil, err
	}

	return &Data{
		Date:             date,
		GitlabUsername:   cfg.GitlabUsername,
		LastXDays:        cfg.LastDays,
		IssueCount:       len(*issues),
		Issues:           *issues,
		ApprovedMRsCount: len(*approvedMRs),
		ApprovedMRs:      *approvedMRs,
		MergedMRsCount:   len(*mergedMRs),
		MergedMRs:        *mergedMRs,
	}, nil
}
