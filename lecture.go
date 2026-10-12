package main

import (
	"fmt"
	"os"
)

func LireFichier(nomFichier string) ([]byte, error) {

	contenu, err := os.ReadFile(nomFichier)

	if err != nil {
		return nil, err
	}

	return contenu, nil
}

func EcrireFichier(nomFichier string, contenu []byte) error {

	err := os.WriteFile(nomFichier, contenu, 0644)

	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture : %w", err)
	}

	return nil
}
