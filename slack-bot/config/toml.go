package config

import (
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
)

var SLACKBOT_CONFIG = os.Getenv("SLACKBOT_CONFIG")

var botConfig *SlackBotConfig

func GetConfig() *SlackBotConfig {
	return botConfig
}

func Init() error {
	var err error
	botConfig, err = ReadConfig(SLACKBOT_CONFIG)
	return err
}

func ReadConfig(path string) (*SlackBotConfig, error) {
	if SLACKBOT_CONFIG == "" {
		log.Fatalf("SLACKBOT_CONFIG environment variable not defined\n")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err;
	}

	decoder := toml.NewDecoder(file)

	config := &SlackBotConfig{}

	if err = decoder.Decode(config); err != nil {
		return nil, err;
	}

	return config, nil
}
