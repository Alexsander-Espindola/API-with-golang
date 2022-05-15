package main

import (
	"fmt"
	"time"
)

func forD(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

func main() {
	forD("sem routines")
	go forD("com routines")
	fmt.Println("Hello mundo")
	fmt.Println("Olá world")
	time.Sleep(time.Second)
	fmt.Println("Terminou")
}
