package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No filename specified, try: go run main.go data")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//File implements Reader interface:
	//https://golang.org/src/os/file.go?s=3908:3956#L102
	//https://golang.org/pkg/io/#Reader

	io.Copy(os.Stdout, file)
}
