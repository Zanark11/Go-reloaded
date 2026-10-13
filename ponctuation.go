package main

import (
	"regexp"
)

func TraiterPonctuation(texte string) string {

	avant := regexp.MustCompile(`[ \t]+([.,!?;:])`)
	texte = avant.ReplaceAllString(texte, "$1")

	apres := regexp.MustCompile(`([.,!?;:])([^\s.,!?;:])`)
	texte = apres.ReplaceAllString(texte, "$1 $2")

	return texte
}
