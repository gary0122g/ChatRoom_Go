package main

import (
	"fmt"
	"net"
)

type room struct {
	member map[net.Addr]*client
	name   string
}

func (r *room) broadcast(sender *client, msg string) {
	for ip, m := range r.member {
		if sender.conn.RemoteAddr() != ip {
			m.msg(fmt.Sprintf("from %s :", sender.nick) + msg)
		}
	}
}
