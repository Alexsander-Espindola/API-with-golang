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

func rodandoForD() {
	forD("sem routines")
	go forD("com routines")
	fmt.Println("Hello mundo")
	fmt.Println("Olá world")
	time.Sleep(time.Second)
	fmt.Println("Terminou")
}

func forTimeSleep(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, "numero:", i)
		time.Sleep(time.Second)
	}
}

func rodandoForTimeSleep() {
	// forTimeSleep("sem routines 1")
	// forTimeSleep("sem routines 2")

	go forTimeSleep("com routines 2")
	go forTimeSleep("com routines 2")

}

func main() {
	rodandoForD()
	rodandoForTimeSleep()
	time.Sleep(time.Second * 10)

}
