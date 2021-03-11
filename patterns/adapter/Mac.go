package main

import "fmt"

type mac struct {
}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("lighting is connected")
}
