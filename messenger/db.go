package messenger

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// mock db message table
var messages = []*Message{
	{ID: "1", UserID: 100, Message: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", Date: time.Unix(1530688169, 0).Add(time.Hour * 1)},
	{ID: "2", UserID: 100, Message: "Gravida dictum fusce ut placerat.", Date: time.Unix(1530688169, 0).Add(time.Hour * 792)},
	{ID: "3", UserID: 100, Message: "Aliquam etiam erat velit scelerisque in. Eget magna fermentum iaculis eu non diam.", Date: time.Unix(1530688169, 0).Add(time.Hour * 360)},
	{ID: "4", UserID: 200, Message: "Orci eu lobortis elementum nibh tellus molestie nunc non blandit.", Date: time.Unix(1530688169, 0)},
	{ID: "5", UserID: 300, Message: "Id donec ultrices tincidunt arcu non. Suspendisse potenti nullam ac tortor vitae.", Date: time.Unix(1530688169, 0).Add(time.Hour * 1)},
	{ID: "6", UserID: 300, Message: "Mi sit amet mauris commodo quis.", Date: time.Unix(1530688169, 0).Add(time.Hour * 100)},
	{ID: "7", UserID: 300, Message: "Lacus sed turpis tincidunt id aliquet.", Date: time.Unix(1530688169, 0).Add(time.Hour * 1000)},
	{ID: "8", UserID: 400, Message: "Mi proin sed libero enim sed.", Date: time.Unix(1530688169, 0)},
	{ID: "9", UserID: 500, Message: "Venenatis urna cursus eget nunc.", Date: time.Unix(1530688169, 0).Add(time.Minute * 30)},
	{ID: "10", UserID: 500, Message: "LOL Latin :)", Date: time.Unix(1530688169, 0).Add(time.Hour * 10)},
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
func dbNewMessage(message *Message, uID int) (string, error) {
	message.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
	message.UserID = uID
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

func dbGetUser(id int) (*User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found.")
}
