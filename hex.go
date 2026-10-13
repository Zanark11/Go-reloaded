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
	lignes := strings.Split(texte, "\n")
	var resultatLignes []string

	for _, ligne := range lignes {
		mots := strings.Fields(ligne)
		var resultat []string

		for i := 0; i < len(mots); i++ {
			if i+1 < len(mots) && strings.HasPrefix(mots[i+1], "(hex)") {
				nombre, err := ConvertirHex(mots[i])
				if err != nil {
					return "", err
				}

				resultat = append(resultat, nombre)

				ponctuation := strings.TrimPrefix(mots[i+1], "(hex)")
				if ponctuation != "" {
					resultat[len(resultat)-1] += ponctuation
				}

				i++
			} else {
				resultat = append(resultat, mots[i])
			}
		}

		resultatLignes = append(resultatLignes, strings.Join(resultat, " "))
	}

	return strings.Join(resultatLignes, "\n"), nil
}
