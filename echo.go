package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func getEchoHandler(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check API key
		// localhost:8080/
		queryAPIKey := r.URL.Query().Get("apiKey")
		log.Println("queryAPIKey:", queryAPIKey)
		if queryAPIKey != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Upgrade initial GET request to a WebSocket
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
