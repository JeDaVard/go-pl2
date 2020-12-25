package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var links = []string{
	"https://google.com",
	"https://golang.org",
	"https://github.com",
	"https://stackoverflow.com",
	"https://nodejs.org",
	"https://better-than-JS.org",
	"https://reactjs.org",
}

var results = make(map[string][]int64)

func main() {
	c := make(chan string)

	for _, link := range links {
		results[link] = []int64{}
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(6 * time.Second)
			dur := checkLink(l, c)
			results[l] = append(results[l], dur)
		}(l)
	}
}

func checkLink(link string, c chan string) int64 {
	start := time.Now()
	_, err := http.Get(link)
	duration := time.Since(start).Milliseconds()
	if err != nil {
		fmt.Println("["+link+"] may be down | "+strconv.FormatInt(duration, 10)+" ms", "| avg ", countAvg(results[link]), " ms")
		c <- link
		return duration
	}
	fmt.Println("["+link+"] ok | "+strconv.FormatInt(duration, 10)+" ms", "| avg ", countAvg(results[link]), " ms")
	c <- link
	return duration
}
