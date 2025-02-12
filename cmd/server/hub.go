package main

import (
	"log"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte, 1024),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) getUserList() []string {
	var users []string
	for c := range h.clients {
		users = append(users, c.user)
	}
	return users
}

func (h *Hub) broadcastUserList() {
	users := h.getUserList()
	usersList, err := newUsersList(users)
	if err != nil {
		log.Println("Error creating user list message:", err)
		return
	}

	select {
	case h.broadcast <- usersList:
	default:
		log.Println("Broadcast channel is full, user list not sent")
	}
}

func (h *Hub) run() {
	for {
		select {
		case msg := <-h.broadcast:
			for c := range h.clients {
				select {
				case c.send <- msg:
				default:
					log.Printf("Client %s disconnecting", c.user)
					delete(h.clients, c)
					close(c.send)
				}
			}
		case client := <-h.register:
			log.Printf("%s has entered the chat\n", client.user)
			h.clients[client] = true
			h.broadcastUserList()
		case client := <-h.unregister:
			log.Printf("%s has exited the chat\n", client.user)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.broadcastUserList()
			}
		}
	}
}
