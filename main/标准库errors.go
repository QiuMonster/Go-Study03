package main

import (
	"errors"
	"fmt"
	// "os"
)

func go_error() {
	// s, err := checknil("")
	s, err := checknil("hello")
	fmt.Println(s, err)
}
func checknil(msg string) (string, error) {
	if msg == "" {
		err := errors.New("字符串不为空...")
		return "", err
	} else {
		return msg, nil
	}
}
