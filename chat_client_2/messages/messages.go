package messages

import (
	"errors"
	"strings"
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"chat_client_2/state"
)

var ProcessMsgCommands = func (cmdList []string, Presentstate *state.State, Client http.Client) error {
	var err error
	// fmt.Printf("The second param is :%vasfjaf", cmdList[1])
	switch cmdList[1] {
	case "send":
		err = ProcessMsgSend(cmdList, Presentstate, Client)
	case "list":
		err = ProcessMsgList(cmdList, Presentstate, Client)
	default:
		err = errors.New("invalid msg command")
	}
	return err
}

func ProcessMsgSend(cmdList []string, Presentstate *state.State, Client http.Client) error {
	finalMsg := strings.Join(cmdList[2:], " ")
	params := url.Values{}
    params.Add("userId", Presentstate.GetUsername())
	params.Add("roomId", Presentstate.GetCurrentRoomId())
	params.Add("text", finalMsg)
	resp, _ := Client.Get("http://localhost:8080/msg/send?"+params.Encode())
	if resp.StatusCode == 200 {
		fmt.Printf("%v\t\t :%v\n", Presentstate.GetDisplayName(), finalMsg)
	}else{
		fmt.Println("couldn't send message")
	}
	return nil
}

func ProcessMsgList(cmdList []string, Presentstate *state.State, Client http.Client) error {
	params := url.Values{}
	params.Add("roomId", Presentstate.GetCurrentRoomId())
	resp, _ := Client.Get("http://localhost:8080/msg/list?"+params.Encode())
	if resp.StatusCode == 200 {
		messages_list := []string{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(bodyBytes, &messages_list)
		fmt.Println("The Recent messages in current room:")
		for _, msg := range messages_list{
			fmt.Println(msg)
		}
	}else{
		fmt.Println("couldn't fetch messages")
	}
	return nil
}