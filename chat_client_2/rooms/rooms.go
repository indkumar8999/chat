package rooms

import (
	"errors"
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"chat_client_2/state"
)

func ProcessRoomCommands(cmdList []string, Presentstate *state.State, Client http.Client) error {
	var err error
	// fmt.Printf("The second param is :%vasfjaf", cmdList[1])
	switch cmdList[1] {
	case "list":
		err = ProcessRoomList(cmdList, Client)
	case "create":
		err = ProcessRoomCreate(cmdList, Presentstate, Client)
	case "join":
		err = ProcessRoomJoin(cmdList, Presentstate)
	case "leave":
		err = ProcessRoomLeave(cmdList, Presentstate)
	case "delete":
		err = ProcessRoomDelete(cmdList, Presentstate, Client)
	default:
		err = errors.New("invalid room command")
	}
	return err
}

func ProcessRoomList(cmdList []string, Client http.Client) error {
	resp, _ := Client.Get("http://localhost:8080/rooms")
	roomlist := []string{}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(bodyBytes, &roomlist)
	fmt.Println("The rooms available:")
	for _, room := range roomlist{
		fmt.Println(room)
	}
	return nil
}

func ProcessRoomCreate(cmdList []string, Presentstate *state.State, Client http.Client) error {
	params := url.Values{}
    params.Add("userId", Presentstate.GetUsername())
	resp, _ := Client.Get("http://localhost:8080/room/create?"+params.Encode())
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Room Created With Id: %v \n", string(bodyBytes))
	return nil
}

func ProcessRoomJoin(cmdList []string, Presentstate *state.State) error {
	if len(cmdList) > 3 {
		return errors.New("not a valid room command")
	}
	roomId := cmdList[2]
	Presentstate.SetCurrentRoomId(roomId)
	return nil
}

func ProcessRoomLeave(cmdList []string, Presentstate *state.State) error {
	fmt.Println(len(cmdList), cmdList)
	if len(cmdList) > 2 {
		return errors.New("not a valid room command")
	}
	if len(Presentstate.GetCurrentRoomId()) > 0 {
		Presentstate.SetCurrentRoomId("")
		return nil
	}
	return errors.New("not in a room now")
}

func ProcessRoomDelete(cmdList []string, Presentstate *state.State, Client http.Client) error {
	if len(cmdList) > 3 {
		return errors.New("not a valid room command")
	}
	roomId := cmdList[2]
	params := url.Values{}
    params.Add("userId", Presentstate.GetUsername())
	params.Add("roomId", roomId)
	resp, _ := Client.Get("http://localhost:8080/room/delete?"+params.Encode())
	if resp.StatusCode == 200 {
		Presentstate.RemoveRoomFromOwned(roomId)
	}
	return nil
}