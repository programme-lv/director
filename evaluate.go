package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func getEvaluateHandler(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryAPIKey := r.URL.Query().Get("apiKey")
		if queryAPIKey != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		// Ticker to send time every second
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for t := range ticker.C {
			msg := fmt.Sprintf("Time: %s", t.String())
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}
