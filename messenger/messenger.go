package messenger

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func ListMessages(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewMessageListResponse(messages)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func SearchMessages(w http.ResponseWriter, r *http.Request) {
	render.RenderList(w, r, NewMessageListResponse(messages))
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value("user").(User)

	data := &MessageRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	message := data.Message

	dbNewMessage(message, u.ID)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewMessageResponse(message))
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	message := r.Context().Value("message").(*Message)

	if err := render.Render(w, r, NewMessageResponse(message)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// DeleteMessage removes an existing Message from mock db
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var err error

	// use context for id info
	message := r.Context().Value("message").(*Message)

	message, err = dbRemoveMessage(message.ID)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewMessageResponse(message))
}

// insert 30 days/100 limit for messages logic here
func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

type UserPayload struct {
	*User
	Role string `json:"role"`
}

func NewUserPayloadResponse(user *User) *UserPayload {
	return &UserPayload{User: user}
}

func (u *UserPayload) Bind(r *http.Request) error {
	return nil
}

func (u *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	u.Role = "messengerUser"
	return nil
}

type MessageRequest struct {
	*Message
	User        *UserPayload `json:"user,omitempty"`
	ProtectedID string       `json:"id"`
}

func (a *MessageRequest) Bind(r *http.Request) error {
	// return err to avoid null pointers
	if a.Message == nil {
		return errors.New("missing required Message fields.")
	}

	a.ProtectedID = ""
	a.Message.Message = strings.ToLower(a.Message.Message)
	return nil
}

type MessageResponse struct {
	*Message
	User *UserPayload `json:"user,omitempty"`
}

func NewMessageResponse(message *Message) *MessageResponse {
	resp := &MessageResponse{Message: message}

	if resp.User == nil {
		if user, _ := dbGetUser(resp.UserID); user != nil {
			resp.User = NewUserPayloadResponse(user)
		}
	}

	return resp
}

func (rd *MessageResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewMessageListResponse(messages []*Message) []render.Renderer {
	list := []render.Renderer{}
	for _, message := range messages {
		list = append(list, NewMessageResponse(message))
	}
	return list
}
