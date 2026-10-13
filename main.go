package main

import (
	"fmt"
	"os"
)

func main() {

	// Vérifie qu'on a bien deux fichiers en argument
	if len(os.Args) != 3 {
		fmt.Println("Utilisation : go run . fichier_entree fichier_sortie")
		return
	}

	// Récupère les noms des fichiers
	fichierEntree := os.Args[1]
	fichierSortie := os.Args[2]

	// Lit le fichier d'entrée
	contenu, err := LireFichier(fichierEntree)
	if err != nil {
		fmt.Println("Erreur lors de la lecture :", err)
		return
	}

	// Transforme le contenu en texte
	texte := string(contenu)

	// Traite les nombres hexadécimaux
	texte, err = TraiterHex(texte)
	if err != nil {
		fmt.Println("Erreur lors de la conversion hexadécimale :", err)
		return
	}

	// Traite les nombres binaires
	texte, err = TraiterBin(texte)
	if err != nil {
		fmt.Println("Erreur lors de la conversion binaire :", err)
		return
	}

	// On traite les changements de casse.
	texte = TraiterCasse(texte)

	// On corrige les espaces autour de la ponctuation.
	texte = TraiterPonctuation(texte)

	// Écrit le texte transformé
	err = EcrireFichier(fichierSortie, []byte(texte))
	if err != nil {
		fmt.Println(err)
		return
	}

}
