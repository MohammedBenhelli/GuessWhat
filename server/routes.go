package server

import (
	"encoding/json"
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
			persons: 1,
			Chat: Chat{Messages: []Message{}},
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

func joinChannel(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONJoinChannel
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else if _, exist := s.channel.Channels[f.RoomName]; !exist {
		resp.Error = "Error room name don't exist !"
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
		s.userList.Users[f.Username] = &u
		if err := s.channel.Channels[f.RoomName].addToTeam(&u); err != nil {
			return err
		}
		resp.Message = "Added to channel"
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
		resp.Message = "update-canvas"
		resp.Data = f.Canvas
		l, err := s.getChannel(&f.RoomName)
		if err != nil {
			return err
		}
		if f.Canvas == l.Canvas {
			return nil
		}
		if err := l.updateCanvas(conn, resp); err != nil {
			return err
		}
	}
	return nil
}

func isAdmin(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONGetRoomName
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else {
		resp.Message = "is-admin"
		l, err := s.getChannel(&f.RoomName)
		if err != nil {
			return err
		}
		if err := l.isAdmin(conn, resp); err != nil {
			return err
		}
	}
	return nil
}

func isDrawer(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONGetRoomName
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else {
		resp.Message = "is-drawer"
		l, err := s.getChannel(&f.RoomName)
		if err != nil {
			return err
		}
		if err := l.isDrawer(conn, resp); err != nil {
			return err
		}
	}
	return nil
}

func startGame(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONGetRoomName
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else {
		resp.Message = "start-game"
		l, err := s.getChannel(&f.RoomName)
		if err != nil {
			return err
		}
		if err := l.startLobby(); err != nil {
			return err
		}
		js, err := json.Marshal(l)
		if err != nil {
			return err
		}
		resp.Data = string(js)
		for i := range l.Teams {
			for j := range l.Teams[i].Users {
				if err := l.Teams[i].Users[j].conn.WriteMessage(1, resp.toJSON()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func newMessage(s *Server, conn *websocket.Conn, p *[]byte) error {
	var f JSONNewMessage
	resp := createJSONResp()
	if err := json.Unmarshal(*p, &f); err != nil {
		resp.Error = "Error can't read request !"
		if err := conn.WriteMessage(1, resp.toJSON()); err != nil {
			return err
		}
	} else {
		resp.Message = "new-message"
		l, err := s.getChannel(&f.RoomName)
		if err != nil {
			return err
		}
		if err := l.newMessage(&f.Message, conn); err != nil {
			return err
		}
		js, err := json.Marshal(l)
		if err != nil {
			return err
		}
		resp.Data = string(js)
		for i := range l.Teams {
			for j := range l.Teams[i].Users {
				if err := l.Teams[i].Users[j].conn.WriteMessage(1, resp.toJSON()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
