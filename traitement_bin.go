
package main

import (
	"strings"
)

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

			i++
		} else {

			resultat = append(resultat, mots[i])
		}
	}

	return strings.Join(resultat, " "), nil
}
