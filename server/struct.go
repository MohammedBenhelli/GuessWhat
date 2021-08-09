package server

import (
	"github.com/pkg/errors"
)

const (
	FILENAME string = "./server/words.txt"
	WORDS    int    = 1799
)

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

func (s *Server) addChannel(f *JSONCreateChannel) error {
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
	}
	return nil
}
