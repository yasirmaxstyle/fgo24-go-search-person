package main

import (
	"fmt"
	"strings"
)

func searchPerson(name string) []string {
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

	found := []string{}

	for x := range users {
		if strings.Contains(strings.ToLower(users[x]), strings.ToLower(name)) {
			found = append(found, users[x])
		}
	}

	return found
}

func main() {
	fmt.Println(searchPerson("clem"))
}
