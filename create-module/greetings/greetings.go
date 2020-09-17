package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// message := fmt.Sprintf("Hello! Welcome %v!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats :[]string{
		"Hi, %v. Welcome!", 
		"Great to see you %v!",
		"哈囉 %v!",
	}

	return formats[rand.Intn(len(formats))]
}