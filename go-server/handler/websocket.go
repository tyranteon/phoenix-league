package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"smaug/phoenix-league/go-server/log"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocket struct {
	db *sql.DB
}

func NewWebSocket(db *sql.DB) *WebSocket {
	return &WebSocket{
		db: db,
	}
}

type msg struct {
	Mutation string `json:"mutation"`
}

func (*WebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println("We got a connection!")

	//w.Write([]byte(`{"mutation":"login"}`))

	time.AfterFunc(time.Second*10, func() {
		conn.WriteJSON(msg{Mutation: "login"})
	})

	_ = conn
}
