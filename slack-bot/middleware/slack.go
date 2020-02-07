package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	slackevents "github.com/nlopes/slack/slackevents"

	"slack-bot/constants"
	"slack-bot/config"
)

func SlackEventParse(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SlackEventParse")

		botConfig := config.GetConfig()

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		log.Println(body)

		event, e := slackevents.ParseEvent(json.RawMessage(body),
				slackevents.OptionVerifyToken(
					&slackevents.TokenComparator {
						VerificationToken: botConfig.SlackToken,
					}))
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, constants.SLACK_EVENT, &event)

		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	}
}
