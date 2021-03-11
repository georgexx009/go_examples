package main

import "fmt"

type windows struct {
}

func (w *windows) connectUsb() {
	fmt.Println("usb is connected")
}
