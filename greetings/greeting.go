package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Hello func - Name parameter cannot be empty")
	}

	randomFormatString := randomFormat()
	// if a function does not return an error, it will return a 
	// single value
	message := fmt.Sprintf(randomFormatString, name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
	// A map (apparently) is like a hash table/dictionary
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)

        if err != nil {
            return nil, err
        }

        messages[name] = message
    }
    return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := [3]string {
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hello there, %v!",
	}

	return formats[rand.Intn(len(formats))]
}