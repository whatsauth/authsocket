package wasocket

import (
	"encoding/json"
	"log"
)

func SendMessageTo(ID string, msg string) (res bool) {
	m := Message{
		Id:      ID,
		Message: msg,
	}
	if Clients[ID] == nil {
		res = false
	} else {
		SendMesssage <- m
		res = true
	}
	return
}

func SendStructTo(ID string, strc interface{}) (res bool) {
	b, err := json.Marshal(strc)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	return SendMessageTo(ID, string(b))
}
