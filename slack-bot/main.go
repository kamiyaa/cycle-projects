package main

import (
	"fmt"
	"log"
	"net/http"

	"slack-bot/bot"
	"slack-bot/config"
	"slack-bot/handlers"
)

func main() {
	const __function_name = "main"

	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	bot.Init()

	botConfig := config.GetConfig()

	for _, h := range handlers.Handlers {
		http.HandleFunc(h.Path, h.HandleFunc)
		log.Printf("[%s] Registered handler for: %s\n",
			__function_name, h.Path)
	}

	log.Printf("[%s] Running server... on port %d\n",
		__function_name,
		botConfig.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", botConfig.Port), nil)
}
