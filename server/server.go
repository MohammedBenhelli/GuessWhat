package server

import "github.com/pkg/errors"

func (s *Server) getChannel(r *string) (*Lobby, error) {
	if _, exist := s.channel.Channels[*r]; exist {
		return nil, errors.New("Can't find the channel !")
	}
	return s.channel.Channels[*r], nil
}
