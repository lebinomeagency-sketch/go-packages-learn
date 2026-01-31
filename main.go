package main

import (
	"fmt"

	"github.com/lebinomeagency-sketch/go-packages-learn/greetings"
)

func main() {
	fmt.Println("Hello World")
	greetings.SayHello("Igor")
	greetings.LogInfo()
}
