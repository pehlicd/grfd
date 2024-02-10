package internal

import "time"

func (client *Client) DataGenerator(cfg Config) (*Data, error) {
	date := time.Now().Format("02.01.2006")
	issues, err := client.getLastXDaysIssues(cfg)
	if err != nil {
		return nil, err
	}

	return &Data{
		Date:       date,
		LastXDays:  cfg.LastDays,
		Issues:     *issues,
		IssueCount: len(*issues),
	}, nil
}
