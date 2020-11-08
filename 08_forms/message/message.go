package message

import (
	"regexp"
	"strings"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

// Message type
type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

// Validate function
func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))
	if match == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Please enter a message"
	}

	return len(msg.Errors) == 0
}
