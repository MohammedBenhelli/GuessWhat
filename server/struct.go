package server

import (
	"github.com/gorilla/websocket"
)

const (
	FILENAME string = "./server/words.txt"
	WORDS    int    = 1799
)

var ROUTER = map[string]RouteHandler{
	"create-channel": addChannel,
	"update-canvas":  updateCanvas,
}

type RouteHandler func(s *Server, conn *websocket.Conn, p *[]byte) error

type User struct {
	conn     *websocket.Conn
	Username string `json:"username"`
	Score    uint   `json:"score"`
}

type Team struct {
	Users []*User `json:"users"`
}

type Message struct {
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
	User      User   `json:"user"`
}

type Chat struct {
	Messages []Message `json:"messages"`
}

type Lobby struct {
	Teams   []Team `json:"teams"`
	Name    string `json:"name"`
	Admin   *User  `json:"admin"`
	Started bool   `json:"started"`
	Canvas  string `json:"canvas"`
	Word    string `json:"word"`
	Drawer  *User  `json:"drawer"`
}

type JSONCreateChannel struct {
	Username string `json:"username"`
	RoomName string `json:"room_name"`
}

type JSONUpdateCanvas struct {
	Canvas   string `json:"canvas"`
	RoomName string `json:"room_name"`
}

type Channel struct {
	Channels map[string]*Lobby `json:"channels"`
}

type UserList struct {
	Users map[string]*User `json:"users"`
}

type Server struct {
	userList UserList
	channel  Channel
}

type RequestType struct {
	Type string `json:"type"`
}

type JSONResp struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
