package main

import "encoding/json"

const (
	EventUserList = iota
	EventMessage
)

type UsersListResponse struct {
	Event int      `json:"event"`
	Users []string `json:"users"`
}

type MessageResponse struct {
	Event    int    `json:"event"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

func newUsersList(users []string) ([]byte, error) {
	msg := &UsersListResponse{
		Event: 0,
		Users: users,
	}
	return json.Marshal(msg)
}

func newMessage() ([]byte, error) {
	msg := &MessageResponse{
		Event:    1,
		Username: "",
		Content:  "",
	}
	return json.Marshal(msg)
}
