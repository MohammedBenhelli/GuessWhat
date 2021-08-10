package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func addChannel(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONCreateChannel
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else if _, exist := s.channel.Channels[f.RoomName]; exist {
		resp.Error = "Error room name already exist !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else if _, exist := s.userList.Users[f.Username]; exist {
		resp.Error = "Error username already exist !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else {
		u := User{conn, f.Username, 0}
		l := Lobby{
			Teams:   []Team{},
			Name:    f.RoomName,
			Admin:   &u,
			Started: false,
		}
		s.userList.Users[f.Username] = &u
		s.channel.Channels[f.RoomName] = &l
		if err := s.channel.Channels[f.RoomName].addToTeam(&u); err != nil {
			return err
		}
		resp.Message = "Channel created"
		resp.Data = f.RoomName
		if err := u.conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
		log.Println(*s.channel.Channels[f.RoomName])
	}
	return nil
}

func updateCanvas(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONUpdateCanvas
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else {
		fmt.Println(s.getChannel(&f.RoomName))
	}
	return nil
}
