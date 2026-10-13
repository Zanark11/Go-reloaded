package main

import (
	"strconv"
	"strings"
)

func TraiterCasse(texte string) string {
	lignes := strings.Split(texte, "\n")
	var resultatLignes []string

	for _, ligne := range lignes {
		mots := strings.Fields(ligne)
		var resultat []string

		for i := 0; i < len(mots); i++ {
			mot := mots[i]

			if (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots) {
				mot += " " + mots[i+1]
				i++
			}

			typeCasse, nombreMots, ponctuation := lireMarqueur(mot)

			if typeCasse == "" {
				resultat = append(resultat, mot)
				continue
			}

			debut := len(resultat) - nombreMots
			if debut < 0 {
				debut = 0
			}

			for j := debut; j < len(resultat); j++ {
				switch typeCasse {
				case "up":
					resultat[j] = strings.ToUpper(resultat[j])
				case "low":
					resultat[j] = strings.ToLower(resultat[j])
				case "cap":
					resultat[j] = capitaliser(resultat[j])
				}
			}

			if len(resultat) > 0 {
				resultat[len(resultat)-1] += ponctuation
			}
		}

		resultatLignes = append(resultatLignes, strings.Join(resultat, " "))
	}

	return strings.Join(resultatLignes, "\n")
}

func lireMarqueur(mot string) (string, int, string) {
	fin := strings.Index(mot, ")")
	if fin == -1 {
		return "", 0, ""
	}

	marqueur := mot[:fin+1]
	ponctuation := mot[fin+1:]

	switch marqueur {
	case "(up)":
		return "up", 1, ponctuation
	case "(low)":
		return "low", 1, ponctuation
	case "(cap)":
		return "cap", 1, ponctuation
	}

	for _, casse := range []string{"up", "low", "cap"} {
		if strings.HasPrefix(marqueur, "("+casse+",") {
			parties := strings.Split(marqueur, ",")
			if len(parties) != 2 {
				return "", 0, ""
			}

			nombre, err := strconv.Atoi(
				strings.TrimSpace(strings.TrimSuffix(parties[1], ")")),
			)
			if err != nil || nombre < 1 {
				return "", 0, ""
			}

			return casse, nombre, ponctuation
		}
	}

	return "", 0, ""
}

func capitaliser(mot string) string {
	runes := []rune(strings.ToLower(mot))
	if len(runes) == 0 {
		return mot
	}

	runes[0] = []rune(strings.ToUpper(string(runes[0])))[0]
	return string(runes)
}
