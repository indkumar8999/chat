package rooms

import (
	"github.com/google/uuid"
	"chat_server/messages"
	"fmt"
)

const (
	//TODO: make sure to implement the concurrent logic before increasing the pool size
	POOL_SIZE = 1
)

type Room struct {
	Id string
	NextSeqId int
	MessagesList []*messages.Message
	OwnerUserId string
	ActiveUserIds map[string]struct{}
	IncomingMessages chan *messages.Message
}

func NewRoom(userId string) *Room {
	id := uuid.New()
	return &Room{
		Id: id.String(),
		NextSeqId: int(0),
		MessagesList: []*messages.Message{},
		OwnerUserId: userId,
		ActiveUserIds: make(map[string]struct{}),
		IncomingMessages: make(chan *messages.Message),
	}
}

func(r *Room) GetNextSeqId() int {
	return r.NextSeqId
}

type RoomManager struct {
	RoomsMap map[string]*Room
	TaskQueue chan *Task
	Quit chan bool
}

type Task struct {
	taskType string
	taskPayload map[string]string
}

func(t *Task) Do(rman *RoomManager) {
	switch t.taskType {
	case "send_message":
		// TODO: implement
		msgText := t.taskPayload["Text"]
		userId := t.taskPayload["UserId"]
		roomId := t.taskPayload["RoomId"]
		// _, present := rman.RoomsMap[roomId].ActiveUserIds[userId]
		// if  !present{
		// 	// TODO: handle error case
		// 	return
		// }
		sid := rman.RoomsMap[roomId].NextSeqId
		newMsg := &messages.Message{
			UserId: userId,
			Text: msgText,
			RoomId: roomId,
			SeqId: sid,
		}
		rman.RoomsMap[roomId].NextSeqId = sid + 1
		rman.RoomsMap[roomId].MessagesList = append(rman.RoomsMap[roomId].MessagesList, newMsg)
		fmt.Println("------------------------------------------------------------")
		fmt.Println(rman.RoomsMap[roomId].MessagesList)
		fmt.Println("------------------------------------------------------------")
	case "delete_room":
		// TODO: implement
		userId := t.taskPayload["UserId"]
		roomId := t.taskPayload["RoomId"]
		fmt.Println(userId, roomId)
		fmt.Printf("%v \n", rman.RoomsMap[roomId])
		if rman.RoomsMap[roomId].OwnerUserId == userId {
			delete(rman.RoomsMap, roomId)
		}
	default:
		
	}
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		RoomsMap: make(map[string]*Room),
		TaskQueue: make(chan *Task, POOL_SIZE),
		Quit: make(chan bool),
	}
}

func (rman *RoomManager) Start() {
	for i:=0 ;i<POOL_SIZE;i=i+1{
		go func() {
			for {
				select {
				case task := <- rman.TaskQueue:
					task.Do(rman)
				case <- rman.Quit:
					return
				}
			}
		}()
	}
}

func (rman *RoomManager) AddTask(t *Task) {
	go func() {
		rman.TaskQueue <- t
	}()
}

func (rman *RoomManager) Stop() {
	for i:=0 ;i<POOL_SIZE;i=i+1{
		go func() {
			rman.Quit <-true
		}()
	}
}