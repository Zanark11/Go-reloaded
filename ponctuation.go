package main

import (
	"regexp"
)

func TraiterPonctuation(texte string) string {

	// On repère les espaces placés avant un signe de ponctuation.
	avant := regexp.MustCompile(`[ \t]+([.,!?;:])`)
	texte = avant.ReplaceAllString(texte, "$1")

	// On ajoute un espace après un signe de ponctuation lorsqu'il est directement suivi d'un autre mot.
	apres := regexp.MustCompile(`([.,!?;:])([^\s.,!?;:])`)
	texte = apres.ReplaceAllString(texte, "$1 $2")

	return texte
}
