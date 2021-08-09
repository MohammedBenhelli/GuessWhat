package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"log"
)

const (
	FILENAME string = "./server/words.txt"
	WORDS    int    = 1799
)

var ROUTER = map[string]func(s *Server, conn *websocket.Conn, p *[]byte) error{
	"create-channel": addChannel,
}

type User struct {
	Username string `json:"username"`
	Score    uint   `json:"score"`
}

type Team struct {
	Users []User `json:"users"`
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
	Admin   User   `json:"admin"`
	Started bool   `json:"started"`
}

type JSONCreateChannel struct {
	Username string `json:"username"`
	RoomName string `json:"room_name"`
}

type Channel struct {
	Channels map[string]Lobby `json:"channels"`
}

type UserList struct {
	Users map[string]User `json:"users"`
}

type Server struct {
	userList UserList
	channel  Channel
}

type RequestType struct {
	Type string `json:"type"`
}

func addChannel(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONCreateChannel
	if err := json.Unmarshal(*p, &f); err != nil {
		if err := conn.WriteMessage(1, []byte("Error cant't read request")); err != nil {
			return err
		}
	} else {
		log.Println(f, string(*p))
	}
	if _, exist := s.channel.Channels[f.RoomName]; exist {
		return errors.New("Room name already used!")
	} else if _, exist := s.userList.Users[f.Username]; exist {
		return errors.New("Username already used!")
	} else {
		s.channel.Channels[f.RoomName] = Lobby{
			Teams:   nil,
			Name:    f.RoomName,
			Admin:   User{f.Username, 0},
			Started: false,
		}
		log.Println(*s)
	}
	return nil
}
