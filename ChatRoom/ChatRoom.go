package chatroom

import "fmt"

type ChatRoom struct {
	Id  string
	Pid []string
	Msg []string
}

func (c *ChatRoom) AddMsg(msg string) {
	c.Msg = append(c.Msg, msg)
}

func (c *ChatRoom) DisplayChat() {
	for _, val := range c.Msg {
		fmt.Println(val)
	}
}