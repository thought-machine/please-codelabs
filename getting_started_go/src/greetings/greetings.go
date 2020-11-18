package greetings

import (
	"math/rand"
)

var greetings = []string{
	"Hello",
	"Bonjour",
	"Marhabaan",
}

func Greeting() string {
	return greetings[rand.Intn(len(greetings))]
}