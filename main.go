package main

import (
	"fmt"
	"os"
)

func main() {

	// Vérifie qu'on a bien donné 2 fichiers en argument
	if len(os.Args) != 3 {
		fmt.Println("Utilisation : go run . fichier_entree fichier_sortie")
		return
	}

	// Récupère le nom du fichier d'entrée
	fichierEntree := os.Args[1]

	// Récupère le nom du fichier de sortie
	fichierSortie := os.Args[2]

	// Lit le contenu du fichier d'entrée
	contenu, err := os.ReadFile(fichierEntree)

	// Vérifie s'il y a eu une erreur
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Écrit le contenu dans le fichier de sortie
	err = os.WriteFile(fichierSortie, contenu, 0644)

	// Vérifie s'il y a eu une erreur lors de l'écriture
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}
}