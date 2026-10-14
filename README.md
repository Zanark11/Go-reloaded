## Présentation

**Go-reloaded** est un outil en ligne de commande développé en **Go**.

Son objectif est de lire un fichier texte, d'appliquer différentes transformations et corrections automatiques, puis d'enregistrer le résultat dans un fichier de sortie.

Le programme est lancé avec deux arguments :

- un fichier d'entrée contenant le texte à modifier (sample.txt)
- un fichier de sortie dans lequel le texte transformé est enregistré (result.txt) 

## Fonctionnalités

Le programme permet de :

- convertir les nombres hexadécimaux avec `(hex)` en nombres décimaux ;
- convertir les nombres binaires avec `(bin)` en nombres décimaux ;
- convertir un mot en majuscules avec `(up)` ;
- convertir plusieurs mots en majuscules avec `(up, nombre)` ;
- convertir un mot en minuscules avec `(low)` ;
- convertir plusieurs mots en minuscules avec `(low, nombre)` ;
- mettre une majuscule initiale à un mot avec `(cap)` ;
- mettre une majuscule initiale à plusieurs mots avec `(cap, nombre)` ;
- corriger les espaces autour des signes de ponctuation ;
- gérer les groupes de ponctuation comme `...` ou `!!` ;
- supprimer les espaces inutiles à l'intérieur des apostrophes ;
- transformer `a` en `an` lorsqu'il est suivi d'un mot commençant par une voyelle ou par `h`.

## Prérequis

Pour utiliser le projet, il faut avoir installé :

- **Go** version **1.27 ou supérieure** ;
- **Git**, si le projet doit être cloné depuis GitHub.

## Installation

Cloner le dépôt :

```bash
git clone https://github.com/Zanark11/Go-reloaded.git
```

Entrer dans le dossier :

```bash
cd Go-reloaded
```

Vérifier l'installation de Go :

```bash
go version
```

Vérifier que le projet compile :

```bash
go build ./...
```

## Lancement

Le programme nécessite deux arguments :

```bash
go run . fichier_entree fichier_sortie
```

Par exemple :

```bash
go run . sample.txt result.txt
```

Le programme lit `sample.txt`, applique les différentes transformations et écrit le résultat dans `result.txt`.

Pour afficher le résultat :

```bash
cat result.txt
```

### Test du programme

```bash
go run . sample.txt result.txt
```

Puis :

```bash
cat result.txt
```

## Architecture du projet

Le projet est organisé en plusieurs fichiers Go, chaque fichier étant consacré à une fonctionnalité particulière.

```text
Go-reloaded/
│
├── apostrophe.go
├── article.go
├── bin.go
├── casse.go
├── go.mod
├── hex.go
├── lecture.go
├── main.go
├── ponctuation.go
├── README.md
├── result.txt
└── sample.txt
```

### Rôle des fichiers

**`main.go`**

Point d'entrée du programme. Il récupère les fichiers indiqués en arguments et appelle les différentes fonctions de traitement dans l'ordre.

**`lecture.go`**

Contient les fonctions permettant de lire le fichier d'entrée et d'écrire le résultat dans le fichier de sortie.

**`hex.go`**

Contient la conversion des nombres hexadécimaux en nombres décimaux.

**`bin.go`**

Contient la conversion des nombres binaires en nombres décimaux.

**`casse.go`**

Gère les transformations `(up)`, `(low)`, `(cap)` ainsi que leurs variantes avec un nombre de mots.

**`ponctuation.go`**

Corrige les espaces autour des signes de ponctuation et gère les groupes de ponctuation.

**`apostrophe.go`**

Corrige les espaces placés à l'intérieur des apostrophes.

**`article.go`**

Transforme `a` en `an` lorsque le mot suivant commence par une voyelle ou par `h`.

**`go.mod`**

Contient les informations du module Go et sa version.

## Technologies utilisées

- **Go**
- **Git**
- Bibliothèque standard de Go