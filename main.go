package main

import (
	"fmt"
	"os"
)

func main() {

	// Vérifie qu'on a bien donné deux fichiers en argument
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

	// Écrit le contenu dans le fichier de sortie
	err = EcrireFichier(fichierSortie, contenu)
	if err != nil {
		fmt.Println(err)
		return
	}
}
