package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var CLIENTS = make(map[string]*websocket.Conn)

var UPGRADER = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)

		return
	}

	id := uuid.New().String()

	CLIENTS[id] = conn

	for {
		ty, msg, err := conn.ReadMessage()
		if err != nil || ty != websocket.TextMessage {
			delete(CLIENTS, id)

			return
		}

		for k, v := range CLIENTS {
			if k == id { continue }

			v.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler);
	http.Handle("/", http.FileServer(http.Dir("../../assets")))

	err := http.ListenAndServe("0.0.0.0:" + os.Getenv("PORT"), nil)

	if err != nil {
		fmt.Println("Server failed to start!")

		return
	}
}