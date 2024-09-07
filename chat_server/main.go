package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"chat_server/rooms"
	"chat_server/users"
)

func setupRoutes(app *fiber.App) {
	app.Get("/rooms", rooms.GetRoomsList)
	app.Get("/room/create", rooms.CreateNewRoom)
	app.Get("/room/delete", rooms.DeleteRoom)
	app.Get("/room/seqid", rooms.GetSequenceIdOfRoom)

	app.Get("/user/register", users.CreateUser)
	app.Get("/user/msg", users.MessagesForUserOfRoom)

	app.Get("/msg/send", rooms.SendMessageInRoom)
	app.Get("/msg/list", rooms.GetRecentHistoryOfRoom)
}

func main() {
	app := fiber.New()

	//app.Use(csrf.New())
	app.Use(logger.New())

	setupRoutes(app)
	rooms.Manager.Start()
	defer rooms.Manager.Stop()
	users.Manager.Start()

	log.Fatal(app.Listen(":8080"))
}