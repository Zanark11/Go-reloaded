package main

import (
	"fmt"
	"os"
)

// LireFichier lit le contenu d'un fichier.
func LireFichier(nomFichier string) ([]byte, error) {

	// Lit le contenu du fichier
	contenu, err := os.ReadFile(nomFichier)

	// Retourne le contenu et l'erreur éventuelle
	if err != nil {
		return nil, err
	}

	return contenu, nil
}

// EcrireFichier écrit le contenu dans un fichier.
func EcrireFichier(nomFichier string, contenu []byte) error {

	// Écrit le contenu dans le fichier
	err := os.WriteFile(nomFichier, contenu, 0644)

	// Affiche une erreur si l'écriture échoue
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture : %w", err)
	}

	return nil
}
