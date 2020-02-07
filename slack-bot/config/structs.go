package config

type SlackBotConfig struct {
	SlackToken string `toml:"slack_token"`
	Port uint32 `toml:"port"`
	CycleAPIKey string `toml:"cycle_api_key"`
	HubID string `toml:"hub_id"`
}
