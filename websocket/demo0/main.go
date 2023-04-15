package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	var (
		err      error
		mux      *http.ServeMux
		server   http.Server
		upgrader websocket.Upgrader

		newline []byte = []byte{'\n'}
		space   []byte = []byte{' '}

		writeWait = 10 * time.Second

		pongWait = 60 * time.Second

		pingPeriod = (pongWait * 9) / 10

		maxMessageSize int64 = 512
	)
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		http.ServeFile(w, r, "websocket/demo0/index.html")
	})
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		var (
			conn     *websocket.Conn
			message  []byte
			messages chan []byte = make(chan []byte)
			ok       bool
			ticket   *time.Ticker
		)
		if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
			log.Println(err)
			return
		}
		conn.SetReadLimit(maxMessageSize)
		conn.SetReadDeadline(time.Now().Add(pongWait))
		conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
		go func() {
			defer conn.Close()
			for {
				if _, message, err = conn.ReadMessage(); err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Println(err)
					}
					break
				}
				message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
				fmt.Println(string(message))
				messages <- message
			}
		}()
		go func() {
			ticket = time.NewTicker(pingPeriod)
			defer func() {
				ticket.Stop()
				conn.Close()
			}()
			for {
				select {
				case message, ok = <-messages:
					fmt.Println(string(message))
					conn.SetWriteDeadline(time.Now().Add(writeWait))
					if !ok {
						conn.WriteMessage(websocket.CloseMessage, []byte{})
						return
					}
					if err = conn.WriteMessage(websocket.TextMessage, message); err != nil {
						log.Println(err)
						return
					}
				case <-ticket.C:
					conn.SetWriteDeadline(time.Now().Add(writeWait))
					if err = conn.WriteMessage(websocket.PingMessage, nil); err != nil {
						return
					}
				}
			}
		}()
	})
	server = http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}
}
