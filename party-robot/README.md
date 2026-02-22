# Robot de Fête

Bienvenue dans « Robot de Fête » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Forfaits

En Go, une application est organisée en packages.
Un package est une collection de fichiers source situés dans le même dossier.
Tous les fichiers source d'un dossier doivent avoir le même nom de package en haut du fichier.
Par convention, les packages sont nommés de la même manière que le dossier dans lequel ils sont situés.

```go
package greeting
```

Go fournit une vaste bibliothèque standard de packages que vous pouvez utiliser dans votre programme en utilisant le mot-clé `import`.
Les packages de la bibliothèque standard sont importés en utilisant leur nom.

```go
package greeting

import "fmt"
```

Un package importé est ensuite adressé avec le nom du package :

```go
world := "World"
fmt.Sprintf("Hello %s", world)
```

Go détermine si un élément peut être appelé par du code dans d'autres packages en fonction de la façon dont il est déclaré.
Pour rendre une fonction, un type, une variable, une constante ou un champ de struct visible de l'extérieur (connu sous le nom d'_exporté_), le nom doit commencer par une lettre majuscule.

```go
package greeting

// Hello est une fonction publique (appelable depuis d'autres packages).
func Hello(name string) string {
    return "Hello " + name
}

// hello est une fonction privée (non appelable depuis d'autres packages).
func hello(name string) string {
    return "Hello " + name
}
```

## Formatage des Chaînes

Go fournit un package intégré appelé `fmt` (package de format) qui offre une variété de fonctions pour manipuler le format de l'entrée et de la sortie.
La fonction la plus couramment utilisée est `Sprintf`, qui utilise des verbes comme `%s` pour interpoler des valeurs dans une chaîne et retourne cette chaîne.

```go
import "fmt"

food := "taco"
fmt.Sprintf("Bring me a %s", food)
// Retourne : Bring me a taco
```

En Go, les valeurs en virgule flottante sont commodément formatées avec les verbes de Sprintf : `%g` (représentation compacte), `%e` (exposant) ou `%f` (non exposant).
Les trois verbes permettent de contrôler la largeur et la position numérique du champ.

```go
import "fmt"

number := 4.3242
fmt.Sprintf("%.2f", number)
// Retourne : 4.32
```

Vous pouvez trouver la liste complète des verbes disponibles dans la [documentation du package de format][fmt-docs].

`fmt` contient d'autres fonctions pour travailler avec les chaînes, telles que `Println` qui imprime simplement les arguments qu'elle reçoit sur la console et `Printf` qui formate l'entrée de la même manière que `Sprintf` avant de l'imprimer.

[fmt-docs]: https://pkg.go.dev/fmt

## Instructions

Il y avait une fois un programmeur excentrique vivant dans une maison étrange aux fenêtres barricadées.
Un jour, il a accepté un travail d'un tableau d'affichage en ligne pour construire un robot de fête. Le robot est censé accueillir les gens et les aider à trouver leurs sièges. La première édition était très technique et montrait le manque d'interaction humaine du programmeur. Une partie de cela s'est également glissée dans l'édition suivante.

## 1. Accueillir un nouveau client à la fête

Implémentez la fonction `Welcome` pour retourner un message de bienvenue en utilisant le nom donné :

```go
Welcome("Christiane")
// => Welcome to my party, Christiane!
```

## 2. Accueillir un nouveau client dont c'est l'anniversaire aujourd'hui

Implémentez la fonction `HappyBirthday` pour retourner un message d'anniversaire en utilisant le nom et l'âge de la personne.
Malheureusement, le programmeur est un peu vaniteux, donc le robot doit démontrer sa connaissance de l'anniversaire de chaque client.

```go
HappyBirthday("Frank", 58)
// => Happy birthday Frank! You are now 58 years old!
```

## 3. Donner des directions

Implémentez la fonction `AssignTable` pour donner des directions.
Elle doit accepter 5 paramètres.

- Le nom du nouveau client à accueillir (`string`)
- Le numéro de table (`int`)
- Le nom du voisin de table (`string`)
- La direction où trouver la table (`string`)
- La distance jusqu'à la table (`float64`)

Le format exact du résultat peut être vu dans l'exemple ci-dessous.

Le robot fournit le numéro de table au format 3 chiffres.
Si le nombre a moins de 3 chiffres, il obtient des zéros de tête supplémentaires pour devenir 3 chiffres (par exemple, 3 devient 003).
Le robot mentionne également la distance jusqu'à la table.
Heureusement, seulement avec une précision limitée à 1 chiffre.

```go
AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
// =>
// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.
```

## Source

### Créé par

- @tehsphinx

### Contribué par

- @oanaOM
- @bobtfish
