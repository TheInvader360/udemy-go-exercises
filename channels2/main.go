package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://github.com",
		"http://not.found",
		"http://theinvader360.com",
		"http://udemy.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		for { //infinite loop
			go checkLink(<-c, c) //continually recheck
		}
	*/

	/*
		//Equivalent to the above infinite loop
		//  watch channel c
		//  whenever a value comes out of it, assign the value to link
		//  then execute the for loop body
		for link := range c {
			go checkLink(link, c)
		}
	*/

	//Add a delay between subsequent checks (use a function literal)
	for l := range c {
		/*
			//bug: l changes over time!
			go func() {
				time.Sleep(5 * time.Second)
				checkLink(l, c)
			}()
		*/
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "unreachable")
		c <- link
		return
	}
	fmt.Println(link, "ok")
	c <- link
}
