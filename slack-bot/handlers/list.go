package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	slack "github.com/nlopes/slack"
	slackevents "github.com/nlopes/slack/slackevents"

	"slack-bot/bot"
	"slack-bot/constants"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello, world")
	w.Write([]byte("hello, world"))
}

func SayHelloInSlack(w http.ResponseWriter, r *http.Request) {
	const __function_name = "SayHelloInSlack"

	log.Println("SayHelloInSlack")

	ctx := r.Context()
	// Get JWT from context through middleware
	event := ctx.Value(constants.SLACK_EVENT)
	if event == nil {
		log.Printf("[%s] no event found\n",
			__function_name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eventsAPIEvent, ok := event.(*slackevents.EventsAPIEvent)
	if !ok {
		log.Printf("[%s] failed to type cast event\n",
			__function_name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	api := bot.GetAPI()
	log.Printf("[%s] event type: %s\n",
		__function_name, eventsAPIEvent.Type)
	switch eventsAPIEvent.Type {
	case slackevents.URLVerification:
		var res *slackevents.ChallengeResponse

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		err := json.Unmarshal([]byte(body), &res)
		if err != nil {
			log.Printf("[%s] failed to decode json: %s\n",
				__function_name, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(res.Challenge))
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			log.Printf("[%s] Sending slack message...\n",
				__function_name)
			api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		default:
			log.Printf("[%s] data type: %s\n",
				__function_name, ev)
		}
	}
}
