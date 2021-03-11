package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main3() {
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
		wg.Add(1)
		//go getWebsite(website)
		go func(website string) {
			getWebsite(website)
			wg.Done()
		}(website)
	}

	wg.Wait()
}

func getWebsite(webSite string) {
	//defer wg.Done()

	if res, err := http.Get(webSite); err != nil {
		fmt.Println(webSite, " is down")
	} else {
		fmt.Printf("[%d] %s is up \n", res.StatusCode, webSite)
	}
}
