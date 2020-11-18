package main

import (
	"fmt"

	"example_module/src/greetings"
)

func main(){
	fmt.Printf("%s, world!\n", greetings.Greeting())
}