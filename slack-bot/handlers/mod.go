package handlers

import (
	"net/http"

	"slack-bot/middleware"
)

const (
	API_URL				string = ""
	HELLO_URL			string = API_URL + "/hello"
	SLACK_URL			string = "/slack"
	SLACK_HELLO_URL			string = SLACK_URL + "/hello"

)

type HttpRouteHandler struct {
	Path string
	HandleFunc func(http.ResponseWriter, *http.Request)
}

// Array of path -> handlers
var Handlers []HttpRouteHandler = []HttpRouteHandler {
	HttpRouteHandler {
		HELLO_URL,
		SayHello,
	},
	HttpRouteHandler {
		SLACK_HELLO_URL,
		middleware.SlackEventParse(
			SayHelloInSlack,
		),
	},
}
