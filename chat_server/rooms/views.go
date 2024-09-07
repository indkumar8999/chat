package rooms

import  (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var Manager = NewRoomManager()

const (
	DEFAULT_LIMIT = 10
)

func GetRoomsList(c *fiber.Ctx) error {
	rooms_list := []string{}
	for k := range Manager.RoomsMap {
		rooms_list = append(rooms_list, k)
	}
	rooms_list_bytes, _ := json.Marshal(rooms_list)
	return c.SendString(string(rooms_list_bytes))
}

func CreateNewRoom(c *fiber.Ctx) error {
	userId := c.Query("userId")
	room := NewRoom(userId)
	Manager.RoomsMap[room.Id] = room
	return c.SendString(room.Id)
}

func DeleteRoom(c *fiber.Ctx) error {
	userId := c.Query("userId")
	roomId := c.Query("roomId")
	m := make(map[string]string)
	m["UserId"] = userId
	m["RoomId"] = roomId
	delete_task := &Task{
		taskType: "delete_room",
		taskPayload: m,
	}
	Manager.AddTask(delete_task)
	return c.SendString("delete_room Task Created")
}

func GetSequenceIdOfRoom(c *fiber.Ctx) error {
	return nil
}

func SendMessageInRoom(c *fiber.Ctx) error {
	userId := c.Query("userId")
	roomId := c.Query("roomId")
	text := c.Query("text")
	m := make(map[string]string)
	m["UserId"] = userId
	m["RoomId"] = roomId
	m["Text"] = text
	msg_send_task := &Task{
		taskType: "send_message",
		taskPayload: m,
	}
	fmt.Println("------------------------------------------------------------")
	fmt.Println(m)
	fmt.Println("------------------------------------------------------------")
	Manager.AddTask(msg_send_task)
	return c.SendString("msg_send_task Task Created")
}

func GetRecentHistoryOfRoom(c *fiber.Ctx) error {
	roomId := c.Query("roomId")
	msg_list := []string{}
	room := Manager.RoomsMap[roomId]
	for index := int(0); index<room.GetNextSeqId(); index=index+1 {
		text := room.MessagesList[index].GetText()
		userId := room.MessagesList[index].GetUserId()
		msg_list = append(msg_list, fmt.Sprintf("%v :: %v", userId, text))
	}
	rooms_list_bytes, _ := json.Marshal(msg_list)
	return c.SendString(string(rooms_list_bytes))
}
