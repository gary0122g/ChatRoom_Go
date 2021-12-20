package main

import (
	"fmt"
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has joined %s", conn.RemoteAddr().String())
	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}
	c.readinput()
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) nick(c *client, arg []string) {
	c.nick = arg[0]
	fmt.Printf("client%v name  has turn to %v", c.conn.RemoteAddr(), arg[1])
}

func (s *server) join(c *client, arg []string) {
	roomName := arg[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:   roomName,
			member: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.member[c.conn.RemoteAddr()] = c
	c.room = r
	fmt.Println(r.member)

}
func (s *server) msg(c *client, arg []string) {
	c.room.broadcast(c, arg[1])
}
func (s *server) quit(c *client) {
	if c.room != nil {
		c.msg("see you next time")
		old_room := s.rooms[c.room.name]
		delete(s.rooms[c.room.name].member, c.conn.RemoteAddr())
		old_room.broadcast(c, fmt.Sprintf("%s has left us", c.nick))

	}
	c.conn.Close()
}
