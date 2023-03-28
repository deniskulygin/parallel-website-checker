package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var links = []string {
		"https://google.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is not reachable :(")
		c <- link
		return
	}

	fmt.Println(link, "is okay :)")
	c <- link
}

