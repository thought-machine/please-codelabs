package main

import (
	"fmt"

	"github.com/example/module/src/greetings"
)

func main() {
	fmt.Printf("%s, world!\n", greetings.Greeting())
}
