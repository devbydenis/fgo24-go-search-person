package utils

import (
	"fmt"
	"strings"
)

func SearchPerson(users []string, query string) {
	for _, user := range users {
		if (strings.Contains(user, query)) {
			// fmt.Println("query", query)
			// fmt.Println("user", user)
			fmt.Println("Found:", user)
			return
		} else {
			fmt.Println("Not found:", []string{})
		} 
	}
}