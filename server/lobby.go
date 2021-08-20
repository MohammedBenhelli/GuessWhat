package server

import (
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"time"
)

func (l *Lobby) startLobby() error {
	if len(l.Teams) < 2 {
		return errors.New("You must be at least 2 to start a game!")
	}
	l.Started = true
	w, err := RandomWord()
	if err != nil {
		return err
	}
	l.Word = w
	d, err := GetDrawer(l)
	if err != nil {
		return err
	}
	l.Drawer = d
	l.Canvas = ""
	return nil
}

func (l *Lobby) addToTeam(u *User) error {
	if len(l.Teams) == 0 {
		l.Teams = append(l.Teams, Team{Users: []*User{u}})
	} else if len(l.Teams) == 1 {
		l.Teams = append(l.Teams, Team{Users: []*User{u}})
	} else if len(l.Teams[0].Users) < len(l.Teams[0].Users) {
		l.Teams[0].Users = append(l.Teams[0].Users, u)
	} else {
		l.Teams[1].Users = append(l.Teams[0].Users, u)
	}
	l.persons++
	return nil
}

func (l *Lobby) updateCanvas(conn *websocket.Conn, r *JSONResp) error {
	for i := range l.Teams {
		for j := range l.Teams[i].Users {
			if l.Teams[i].Users[j].conn != conn {
				if err := l.Teams[i].Users[j].conn.WriteMessage(1, r.toJSON()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (l *Lobby) isAdmin(conn *websocket.Conn, r *JSONResp) error {
	if l.Admin.conn == conn {
		r.Data = "true"
		if err := conn.WriteMessage(1, r.toJSON()); err != nil {
			return err
		}
		return nil
	}
	r.Data = "false"
	if err := conn.WriteMessage(1, r.toJSON()); err != nil {
		return err
	}
	return nil
}

func (l *Lobby) isDrawer(conn *websocket.Conn, r *JSONResp) error {
	if l.Drawer.conn == conn {
		r.Data = "true"
		if err := conn.WriteMessage(1, r.toJSON()); err != nil {
			return err
		}
		return nil
	}
	r.Data = "false"
	if err := conn.WriteMessage(1, r.toJSON()); err != nil {
		return err
	}
	return nil
}

func (l *Lobby) newMessage(msg *string, conn *websocket.Conn) error {
	var m Message
	for i := range l.Teams {
		for j := range l.Teams[i].Users {
			if l.Teams[i].Users[j].conn == conn {
				m.User = l.Teams[i].Users[j]
				m.Timestamp = time.Now().String()
				m.Text = *msg
			}
		}
	}
	if l.Word == *msg && l.Drawer.conn != conn {
		m.User.Score += 300
		m.Text = m.User.Username + " a devine le mot !"
		if err := l.startLobby(); err != nil {
			return err
		}
	}
	l.Chat.Messages = append(l.Chat.Messages, m)
	return nil
}
