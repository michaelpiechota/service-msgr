package messenger

import (
	"errors"
	"fmt"
	"math/rand"
)

// mock db message table
var messages = []*Message{
	{ID: "1", UserID: 100, Message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
	{ID: "2", UserID: 100, Message: "Gravida dictum fusce ut placerat."},
	{ID: "3", UserID: 100, Message: "Aliquam etiam erat velit scelerisque in. Eget magna fermentum iaculis eu non diam."},
	{ID: "4", UserID: 200, Message: "Orci eu lobortis elementum nibh tellus molestie nunc non blandit."},
	{ID: "5", UserID: 300, Message: "Id donec ultrices tincidunt arcu non. Suspendisse potenti nullam ac tortor vitae."},
	{ID: "6", UserID: 300, Message: "Mi sit amet mauris commodo quis."},
	{ID: "7", UserID: 300, Message: "Lacus sed turpis tincidunt id aliquet."},
	{ID: "8", UserID: 400, Message: "Mi proin sed libero enim sed."},
	{ID: "9", UserID: 500, Message: "Venenatis urna cursus eget nunc."},
	{ID: "10", UserID: 500, Message: "LOL Latin :)"},
}

// mock db user table
var users = []*User{
	{ID: 100, Name: "Michael"},
	{ID: 200, Name: "Madeline"},
	{ID: 300, Name: "Izzy"},
	{ID: 400, Name: "Honey"},
	{ID: 500, Name: "Aneta"},
}

// *************** helper funcs for the mock db ******************
func dbNewMessage(message *Message) (string, error) {
	message.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
	messages = append(messages, message)
	return message.ID, nil
}

func dbGetMessage(id string) (*Message, error) {
	for _, a := range messages {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("message not found.")
}

func dbRemoveMessage(id string) (*Message, error) {
	for i, a := range messages {
		if a.ID == id {
			messages = append((messages)[:i], (messages)[i+1:]...)
			return a, nil
		}
	}
	return nil, errors.New("message not found.")
}

func dbGetUser(id int64) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}
