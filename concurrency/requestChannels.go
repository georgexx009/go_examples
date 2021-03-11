package main

import (
	"fmt"
	"net/http"
)

func main4() {
	c := make(chan string)

	websites := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, website := range websites {
		go getWebsiteWithChannel(website, c)
	}

	// iterating over the range of channel
	// so keeps receiving messages until channel is closed
	for msg := range c {
		fmt.Println(msg)
	}

	// alternate syntax
	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
}

func getWebsiteWithChannel(webSite string, c chan string) {
	if _, err := http.Get(webSite); err != nil {
		c <- webSite + "is down"
	} else {
		c <- webSite + "is up"
	}

}
