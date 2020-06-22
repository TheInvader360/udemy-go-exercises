package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	delete(colors, "green")

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Println(key, val)
	}
}
