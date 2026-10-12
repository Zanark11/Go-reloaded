package main

import (
	"strconv"
	"strings"
)

func ConvertirHex(nombre string) (string, error) {

	decimal, err := strconv.ParseInt(nombre, 16, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(decimal, 10), nil
}

func TraiterHex(texte string) (string, error) {

	mots := strings.Fields(texte)

	var resultat []string

	for i := 0; i < len(mots); i++ {

		if i+1 < len(mots) && strings.HasPrefix(mots[i+1], "(hex)") {

			nombre, err := ConvertirHex(mots[i])
			if err != nil {
				return "", err
			}

			resultat = append(resultat, nombre)

			// je conserve la ponctuation apres le marqueur
			ponctuation := strings.TrimPrefix(mots[i+1], "(hex)")
			if ponctuation != "" {
				resultat[len(resultat)-1] += ponctuation
			}

			// je saute le marqueur (hex)
			i++
		} else {
			// je conserve le mot tel quel
			resultat = append(resultat, mots[i])
		}
	}

	// je reconstitue le texte
	return strings.Join(resultat, " "), nil
}
