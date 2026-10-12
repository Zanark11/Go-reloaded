
package main

import (
	"strings"
)

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

			ponctuation := strings.TrimPrefix(mots[i+1], "(hex)")
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
