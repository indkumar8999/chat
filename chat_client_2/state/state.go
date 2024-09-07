package state

type State struct {
	username string
	displayName string
	currentRoomId string
	ownedRooms map[string]struct{}
}

func NewState(username, displayName string) *State {
	return &State{
		username : username,
		displayName : displayName,
		currentRoomId : "",
		ownedRooms : make(map[string]struct{}, 0),
	}
}

func (s *State) GetUsername() string {
	return s.username
}

func (s *State) GetDisplayName() string {
	return s.displayName
}

func(s *State) GetCurrentRoomId() string {
	return s.currentRoomId
}

func(s *State) SetCurrentRoomId(room_id string) {
	s.currentRoomId = room_id
}

func (s *State) IsMyRoom(room_id string) bool {
	_, ok := s.ownedRooms[room_id]
	return ok
}

func (s *State) AddRoomInOwned(room_id string) {
	s.ownedRooms[room_id] = struct{}{}
}

func (s *State) RemoveRoomFromOwned(room_id string) {
	delete(s.ownedRooms, room_id)
}