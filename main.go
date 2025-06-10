package main

import (
	"fmt"
	"search-person/utils"
)

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
	fmt.Println(utils.SearchPerson("clem", users))
}
