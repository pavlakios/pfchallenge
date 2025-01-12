package main

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Sessions open sessions simultaneously.
func Sessions() {
	wsURL := "ws://localhost:8080/goapp/ws"
	numConnections, err := strconv.Atoi(os.Args[3])
	if err != nil || numConnections <= 0 {
		log.Fatalf("Invalid number of connections: %s", os.Args[3])
	}

	log.Printf("Starting %d WebSocket connections to %s...", numConnections, wsURL)

	var wg sync.WaitGroup
	wg.Add(numConnections)

	for i := 0; i < numConnections; i++ {
		go func(connID int) {
			defer wg.Done()
			u, err := url.Parse(wsURL)
			if err != nil {
				log.Printf("[conn #%d]: Invalid URL - %v", connID, err)
				return
			}

			conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Printf("[conn #%d]: Failed to connect - %v", connID, err)
				return
			}
			defer conn.Close()

			// Keep connection open and send periodic pings
			// Send pings every 1 seconds
			ticker := time.NewTicker(1 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
						log.Printf("[conn #%d]: Failed to send ping - %v", connID, err)
						return
					}

					_, message, err := conn.ReadMessage()
					if err != nil {
						log.Printf("[conn #%d]: error %v", connID, err)
						return
					}
					log.Printf("[conn #%d]: %s", connID, string(message))
				}
			}
		}(i)
	}

	wg.Wait()
	log.Println("All WebSocket connections completed.")

	return
}
