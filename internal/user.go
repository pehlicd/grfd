package internal

func (client *Client) GetCurrentUser(cfg Config) (Config, error) {
	user, _, err := client.Users.CurrentUser()
	if err != nil {
		return cfg, err
	}

	cfg.GitlabUsername = user.Username
	cfg.GitlabUserId = user.ID

	return cfg, nil
}
