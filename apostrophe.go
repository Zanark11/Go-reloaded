package main

import "strings"

func TraiterApostrophes(texte string) string {
	parties := strings.Split(texte, "'")

	for i := 1; i+1 < len(parties); i += 2 {
		parties[i] = strings.Trim(parties[i], " \t")
	}

	return strings.Join(parties, "'")
}
