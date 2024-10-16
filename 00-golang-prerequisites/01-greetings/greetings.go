package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	msg := fmt.Sprintf(randomFormat(), name)
	return msg, nil
}

func Hellos(names[]string) (map[string]string, error){
	msgs := make(map[string]string)

	for _, name := range names{
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		msgs[name] = msg
	}
	return msgs, nil
}

func randomFormat() string {
	formats := []string {
		"Hi %v. Welcome!",
		"Great to see you, %v",
		"Hail %v! Well met!",
		"Wassup %v?!",
	}
	return formats[rand.Intn(len(formats))]
}