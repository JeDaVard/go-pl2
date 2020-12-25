package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://golang.org",
		"https://github.com",
		"https://stackoverflow.com",
		"https://nodejs.org",
		"https://better-than-JS.org",
		"https://reactjs.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	start := time.Now()
	_, err := http.Get(link)
	duration := time.Since(start).Milliseconds()
	c <- "message to chan"
	if err != nil {
		fmt.Println(link, "may be down |", duration, "ms")
		return
	}
	fmt.Println(link, "is ok |", duration, "ms")
}
