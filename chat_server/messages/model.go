package messages

type Message struct {
	UserId string
	Text string
	SeqId int
	RoomId string
}


func (m *Message) GetUserId() string {
	return m.UserId
}

func (m *Message) GetText() string {
	return m.Text
}