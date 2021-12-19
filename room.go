package main

import "net"

type room struct {
	member map[net.Addr]*client
	name   string
}

func (r *room) broadcast(sender *client, msg string) {
	for ip, m := range r.member {
		if sender.conn.RemoteAddr() != ip {
			m.msg(msg)
		}
	}

}
