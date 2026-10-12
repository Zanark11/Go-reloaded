package main

import (
	"strconv"
	"strings"
)

func ConvertirBin(nombre string) (string, error) {

	decimal, err := strconv.ParseInt(nombre, 2, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(decimal, 10), nil
}

func TraiterBin(texte string) (string, error) {

	mots := strings.Fields(texte)

	var resultat []string

	for i := 0; i < len(mots); i++ {

		if i+1 < len(mots) && strings.HasPrefix(mots[i+1], "(bin)") {

			nombre, err := ConvertirBin(mots[i])
			if err != nil {
				return "", err
			}

			resultat = append(resultat, nombre)

			ponctuation := strings.TrimPrefix(mots[i+1], "(bin)")
			if ponctuation != "" {
				resultat[len(resultat)-1] += ponctuation
			}

			// ici on saute le marqueur (bin)
			i++
		} else {
			// je conserve le mot tel quel
			resultat = append(resultat, mots[i])
		}
	}

	// je reconstitue le texte
	return strings.Join(resultat, " "), nil
}
