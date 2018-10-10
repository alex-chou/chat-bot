package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nlopes/slack/slackevents"
)

// Slack handles slack events.
func (s *Server) Slack(w http.ResponseWriter, r *http.Request) {
	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	log.Print(body.String())
	event, err := slackevents.ParseEvent(json.RawMessage(body.String()), slackevents.OptionVerifyToken(&slackevents.TokenComparator{s.Token}))
	if err != nil {
		log.Printf("Error parsing body to event: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch event.Type {
	case slackevents.URLVerification:
		var resp *slackevents.ChallengeResponse
		if err := json.Unmarshal(body.Bytes(), &resp); err != nil {
			log.Printf("Error unmarshalling body to ChallengeResponse: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(resp.Challenge))
		return
	default:
		http.Error(w, fmt.Sprintf("Event type not handled: %s", event.Type), http.StatusInternalServerError)
	}
}
