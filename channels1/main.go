package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://github.com",
		"http://not.found",
		"http://theinvader360.com",
		"http://udemy.com",
	}

	/*
		//sequential synchronous calls (blocking calls cause delay)
		for _, link := range links {
			checkLink(link)
		}
	*/

	/*
		//concurrent threads (main exits before goroutines are done)
		for _, link := range links {
			go checkLink(link)
		}
	*/

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

/*
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "unreachable")
		return
	}
	fmt.Println(link, "ok")
}
*/

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "unreachable")
		c <- "down?"
		return
	}
	fmt.Println(link, "ok")
	c <- "up!"
}
