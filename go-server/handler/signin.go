package handler

import (
	"database/sql"
	"net/http"
	"smaug/phoenix-league/go-server/log"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/solovev/steam_go"
)

type Signin struct {
	db *sql.DB
}

func NewSignin(db *sql.DB) *Signin {
	return &Signin{
		db: db,
	}
}

func (s *Signin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	opId := steam_go.NewOpenId(r)

	if opId.Mode() == "" {
		http.Redirect(w, r, opId.AuthUrl(), 301)
		return
	}

	if opId.Mode() == "cancel" {
		w.Write([]byte("Authorization cancelled"))
		return
	}

	steamId, err := opId.ValidateAndGetId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sid, err := s.getSessionID(steamId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Error(err)
		return
	}

	expire := time.Now().AddDate(1, 0, 0)

	http.SetCookie(w, &http.Cookie{
		Expires: expire,
		Name:    "session",
		Value:   sid,
	})

	http.Redirect(w, r, "http://www.phoenix-league.net", http.StatusFound)
}

func getExistingSessionID(db *sql.DB, steamID string) (string, error) {
	row := db.QueryRow("SELECT session_id FROM users WHERE steam_id=$1", steamID)

	var sessionID string

	if err := row.Scan(&sessionID); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", errors.WithStack(err)
		}
	}
	return sessionID, nil
}

func createSessionID(db *sql.DB, steamID string) (string, error) {
	sidPostfix := "_" + uuid.NewV4().String()

	row := db.QueryRow("INSERT INTO users (steam_id, display_name, session_id) SELECT $1, $2, last_value || $3 FROM users_user_id_seq RETURNING session_id", steamID, "No Name yet", sidPostfix)

	var s string
	if err := row.Scan(&s); err != nil {
		return "", errors.WithStack(err)
	}

	return s, nil
}

func (s *Signin) getSessionID(steamID string) (string, error) {
	sid, err := getExistingSessionID(s.db, steamID)
	if err != nil {
		return "", err
	}
	if sid != "" {
		return sid, nil
	}

	return createSessionID(s.db, steamID)
}
