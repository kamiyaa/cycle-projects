package bot

import (
	slack "github.com/nlopes/slack"

	"slack-bot/config"
)

var api *slack.Client

func GetAPI() *slack.Client {
	return api
}

func Init() {
	botConfig := config.GetConfig()

	api = slack.New(botConfig.Token, slack.OptionDebug(true))
}
