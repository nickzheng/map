package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	for key, value := range colors {
		fmt.Println(key, value)
	}

	fmt.Printf("%+v\n", colors)
	color2 := make(map[string]string)
	color2["black"] = "#ffffff"
	fmt.Printf("%+v\n", color2)
	delete(color2, "black")
	fmt.Printf("%+v\n", color2)
}
