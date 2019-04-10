package main

import (
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

var clients = make([]*websocket.Conn, 0)

func sendToAllWsConnections(data string) {
	for _, v := range clients {
		v.Write([]byte(data))
	}
}

func webSocketHandler() http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		clients = append(clients, ws)
		io.Copy(ws, ws)
	})
}
