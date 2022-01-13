package engine

import (
	"fmt"
	"strings"
)

type IHandler interface {
	Post(cmd Command)
}

type Handler struct {
	eventLoop *EventLoop
}

type Command interface {
	Execute(handler IHandler)
}

type PrintCmd struct {
	Msg string
}

func (pc PrintCmd) Execute(handler IHandler) {
	fmt.Println(pc.Msg)
}

type ReverseCmd struct {
	Str    string

}

func reverseString(str string) string{
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
	   byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
 }

func (dc ReverseCmd) Execute(handler IHandler) {
	r := reverseString(string(dc.Str))
	var printCmd Command = &PrintCmd{r}
	handler.Post(printCmd)
}

type EventLoop struct {
	queue []Command
}

func (eventLoop *EventLoop) Post(cmd Command) {
	eventLoop.queue = append(eventLoop.queue, cmd)
}

func (eventLoop *EventLoop) Start() {
	eventLoop.queue = make([]Command, 0)
}

func (eventLoop *EventLoop) AwaitFinish() {
	for len(eventLoop.queue) > 0 {
		cmd := eventLoop.queue[0]
		eventLoop.queue = eventLoop.queue[1:]
		cmd.Execute(eventLoop)
	}
}

func Parse(str string) Command {
	array := strings.Fields(str)
	if len(array) == 0 {
		return &PrintCmd{Msg: "Error: Empty"}
	}
	if array[0] == "reverse" {
		return &ReverseCmd{Str: array[1]}
			} else if array[0] == "print" {
		if len(array) == 2 {
			return &PrintCmd{Msg: array[1]}
		}
		return &PrintCmd{Msg: "Error: [0] - command  [1] - string"}
	}
	return &PrintCmd{Msg: "Error: There is no such command"}
}
