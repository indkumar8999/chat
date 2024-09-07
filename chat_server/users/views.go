package users

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)
var Manager = NewUserManager()
func CreateUser(c *fiber.Ctx) error {
	userId := c.Query("userId")
	_, ok := Manager.UserMap[userId]
	fmt.Println(Manager.UserMap, ok)
	if !ok {
		Manager.AddUser(userId)
		return c.SendStatus(fiber.StatusOK)
	}
	return c.SendStatus(fiber.StatusNotAcceptable)
}

func MessagesForUserOfRoom(c *fiber.Ctx) error {
	return nil
}