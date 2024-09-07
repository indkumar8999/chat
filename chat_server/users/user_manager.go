package users

import (
	"chat_server/messages"
)

type User struct {
	UserId string
	EventChannel chan *messages.Message
}

type UserManager struct {
	UserMap map[string]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		UserMap: make(map[string]*User),
	}
}

func (uman *UserManager) Start() {

}

func (uman *UserManager) DoesUserNameExists(userId string) bool {
	_, ok := uman.UserMap[userId]
	return !ok
}

func (uman *UserManager) AddUser(userId string) {
	uman.UserMap[userId] = &User{
		UserId: userId,
		EventChannel: make(chan *messages.Message),
	}
}