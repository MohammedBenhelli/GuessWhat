package server

import "github.com/pkg/errors"

func (l *Lobby) startLobby() error {
	if len(l.Teams) < 2 {
		return errors.New("You must be at least 2 to start a game!")
	}
	l.Started = false
	return nil
}

func (l *Lobby) addToTeam(u *User) error {
	if len(l.Teams) == 0{
		l.Teams = append(l.Teams, Team{Users: []*User{u}})
	} else if len(l.Teams) == 1 {
		l.Teams = append(l.Teams, Team{Users: []*User{u}})
	} else if len(l.Teams[0].Users) < len(l.Teams[0].Users) {
		l.Teams[0].Users = append(l.Teams[0].Users, u)
	} else {
		l.Teams[1].Users = append(l.Teams[0].Users, u)
	}
	return nil
}
