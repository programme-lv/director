package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	log.Printf("API Key: \"%v\"", apiKey)

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// Check API key
		queryAPIKey := r.URL.Query().Get("apiKey")
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
	})

	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
