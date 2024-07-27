// concurrent3
// Make the tests pass!

package main_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestSendAndReceive(t *testing.T) {
	var buf bytes.Buffer

	messages := make(chan string)
	sendAndReceive(&buf, messages)

	got := buf.String()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func sendAndReceive(buf *bytes.Buffer, messages chan string) {
	go func() {
		messages <- "Hello"
		messages <- "World"
		close(messages)
	}()

	var greetings []string
	for message := range messages {
		greetings = append(greetings, message)
	}

	fmt.Fprint(buf, strings.Join(greetings, " "))
}
