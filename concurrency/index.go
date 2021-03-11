package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// receive only
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
		fmt.Println("waiting to send")
	}
}

// send only
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		fmt.Println("waiting to receive")
		time.Sleep(time.Second * 3)
	}
}

func main2() {
	// for i := 0; i < 3; i++ {
	// 	go f(i)
	// }
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
