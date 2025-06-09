package main

import (
	"search-person/utils"
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
	
	utils.SearchPerson(users, "Ervin")
}