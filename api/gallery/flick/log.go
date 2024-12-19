package flick

import (
	"fmt"
	"time"
)

var messages []string = []string{}

func timeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func messagesInit() {
	messages = []string{}
}

func messageLog(message string) {
	line := fmt.Sprintf("[%s] %s", timeFormat(time.Now()), message)
	messages = append(messages, line)
}

func getMessages() []string {
	return messages
}
