package main

import "strings"

func TraiterArticles(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		mots := strings.Fields(ligne)

		for j := 0; j < len(mots)-1; j++ {
			if mots[j] == "a" || mots[j] == "A" {
				motSuivant := strings.ToLower(mots[j+1])

				if strings.ContainsRune("aeiouh", rune(motSuivant[0])) {
					if mots[j] == "A" {
						mots[j] = "An"
					} else {
						mots[j] = "an"
					}
				}
			}
		}

		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}
