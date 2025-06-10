package utils

import "strings"

func SearchPerson(name string, users []string) []string {

	found := []string{}

	for x := range users {
		if strings.Contains(strings.ToLower(users[x]), strings.ToLower(name)) {
			found = append(found, users[x])
		}
	}

	return found
}
