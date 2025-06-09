package main

import (
	"fmt"
	"strings"
)

func searchPerson(name string, users []string) []string {

	found := []string{}

	for x := range users {
		if strings.Contains(strings.ToLower(users[x]), strings.ToLower(name)) {
			found = append(found, users[x])
		}
	}

	return found
}

func main() {
	users := []string{
		"Leanne Graham",
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
	fmt.Println(searchPerson("clem", users))
}
