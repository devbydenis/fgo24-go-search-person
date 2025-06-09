package main

import (
	"fmt"
)

func main() {
	users:= []string{
		"Leane Graham",
		"Ervin Howell",
		"Clementine Bauch",
		"Patricia Lebsack",
		"Chelsey Dietrich",
		"Mrs. Dennis Schulist",
		"Kurtis Weissnat",
		"Nicholas Runolfsdottir V",
		"Glenna Reichert",
		"Clementina DuBuque",
	}
	
	searchPerson(users, "Ervin Howell")
}

func searchPerson(users []string, query string) {
	for _, user := range users {
		if (user == query) {
			fmt.Println("Found:", user)
		} else {
			fmt.Println("Not found:", []string{})
			break
		}
	}
}