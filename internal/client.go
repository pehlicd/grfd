package internal

import (
	"crypto/tls"
	"github.com/xanzy/go-gitlab"
	"net/http"
)

func NewClient(cfg Config) (*Client, error) {
	httpClient := &http.Client{}
	if cfg.Insecure {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client, err := gitlab.NewClient(cfg.GitLabToken, gitlab.WithBaseURL(cfg.GitLabAddr), gitlab.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
