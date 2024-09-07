package main

import (
	"bufio"
    "fmt"
    "os"
	"strings"
	"errors"
	"net/http"
	"net/url"

	"chat_client_2/rooms"
	"chat_client_2/messages"
	"chat_client_2/state"

	"github.com/golang/glog"
)

var Presentstate *state.State
var Client = http.Client{}

const (
	maxCount = 5
)

func IsValidUserName(username string) bool {
	params := url.Values{}
    params.Add("userId", username)
	resp, _ := Client.Get("http://localhost:8080/user/register?"+params.Encode())
	return resp.StatusCode == 200
}

func ProcessWhoAmI(cmdList []string) error {
	if len(cmdList) == 1 {
		user := Presentstate.GetUsername()
		fmt.Printf("User is : %v \n", user)
		fmt.Printf("Room is : %v \n", Presentstate.GetCurrentRoomId())
		return nil
	}
	return errors.New("not valid command")
}

func Split(r rune) bool {
    return r == '\n' || r == ' '
}

func ProcessCommand(cmd string) error {
	var err error
	cmdList := strings.FieldsFunc(cmd, Split)
	// fmt.Printf("command is: %v \n", cmdList)
	if len(cmdList) == 0 {
		fmt.Println("")
		return nil
	}
	switch cmdList[0] {
	case "room":
		err = rooms.ProcessRoomCommands(cmdList, Presentstate, Client)
	case "msg":
		err = messages.ProcessMsgCommands(cmdList, Presentstate, Client)
	case "whoami":
		err = ProcessWhoAmI(cmdList)
	default:
		err = fmt.Errorf("not a valid command: %v", cmd)
	}
	return err
}

func main() {
	defer glog.Flush()
	reader := bufio.NewReader(os.Stdin)
	count := 0
	for {
		fmt.Println("Enter username")
		usernameUntrimmed, _ := reader.ReadString('\n')
		username := strings.TrimSpace(usernameUntrimmed)
		if IsValidUserName(username) {
			fmt.Println("Enter displayName")
			displayNameUntrimmed, _ := reader.ReadString('\n')
			displayName := strings.TrimSpace(displayNameUntrimmed)
			Presentstate = state.NewState(username, displayName)
			break
		}else{
			fmt.Println("Username is already in use: Please enter another")
		}
		count += 1
		if count == maxCount {
			fmt.Println("Enough tries: try to be unique")
			break
		}
	}
	fmt.Println("------------------------------------------------------------------")
	for {
		text, _ := reader.ReadString('\n')
		err := ProcessCommand(text)
		if err != nil {
			fmt.Printf("Error: %v \n", err.Error())
		}
		fmt.Println("------------------------------------------------------------------")
	}
}