package greetings

import (
	"fmt"
	"log"
)

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func secretLog() {
	log.Println("ошибка")
}

func LogInfo() {
	secretLog()
}

func SayGoodye() {
	fmt.Println("Goodbye")
}
