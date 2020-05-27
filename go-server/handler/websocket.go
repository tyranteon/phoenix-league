package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"smaug/phoenix-league/go-server/log"

	"github.com/gorilla/websocket"
	"github.com/sanity-io/litter"
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

type request struct {
	Kind string `json:"kind"`
}

func (s *WebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer conn.Close()

	var uid int
	c, err := r.Cookie("session")
	if err != http.ErrNoCookie {
		if err != nil {
			log.Error(err)
			return
		}
		row := s.db.QueryRow("SELECT user_id FROM users WHERE session_id=$1", c.Value)
		if err := row.Scan(&uid); err != nil {
			log.Error(err)
			return
		}
		conn.WriteJSON(msg{Mutation: "login"})
	} else {
		fmt.Println("Couldn't find the cookie :(")
	}

	fmt.Println("We got a connection!")

	for {
		var req request
		if err := conn.ReadJSON(&req); err != nil {
			if !websocket.IsCloseError(err, websocket.CloseGoingAway) {
				log.Error(err)
			}
			break
		}

		litter.Dump(req)
	}
}
