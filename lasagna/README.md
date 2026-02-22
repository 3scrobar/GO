# Les magnifiques lasagnes de Gopher

Bienvenue dans « Les magnifiques lasagnes de Gopher » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

[Go](https://golang.org) est un langage de programmation compilé et typé statiquement.
Cet exercice présente trois fonctionnalités majeures du langage : packages, fonctions et variables.

## Forfaits

Les applications Go sont organisées en packages.
Un package est un ensemble de fichiers sources situés dans le même répertoire.
Tous les fichiers sources d'un répertoire doivent partager le même nom de package.
Lorsqu'un package est importé, seules les entités (fonctions, types, variables, constantes) dont les noms commencent par une majuscule peuvent être utilisées/accédées.
Le style de dénomination recommandé dans Go est que les identifiants seront nommés en utilisant `camelCase`, à l'exception de ceux destinés à être accessibles entre les packages qui devraient être `PascalCase`.

```go
package lasagna
```

##Variables

Go est typé statiquement, ce qui signifie que toutes les variables [must have a defined type](https://en.wikipedia.org/wiki/Type_system) au moment de la compilation.

Les variables peuvent être définies en spécifiant explicitement un type :

```go
var explicit int // Explicitly typed
```

Vous pouvez également utiliser un initialiseur et le compilateur attribuera le type de variable pour qu'il corresponde au type de l'initialiseur.

```go
implicit := 10   // Implicitly typed as an int
```

Une fois déclarées, les variables peuvent se voir attribuer des valeurs à l'aide de l'opérateur `=`.
Une fois déclaré, le type d'une variable ne peut jamais changer.

```go
count := 1 // Assign initial value
count = 2  // Update to new value

count = false // This throws a compiler error due to assigning a non `int` type
```

## Constantes

Les constantes contiennent une donnée tout comme les variables, mais leur valeur ne peut pas changer pendant l'exécution du programme.

Les constantes sont définies à l'aide du mot-clé `const` et peuvent être des nombres, des caractères, des chaînes ou des booléens :

```go
const Age = 21 // Defines a numeric constant 'Age' with the value of 21
```

## Fonctions

Les fonctions Go acceptent zéro ou plusieurs paramètres.
Les paramètres doivent être explicitement typés, il n’y a pas d’inférence de type.

Les valeurs sont renvoyées par les fonctions à l'aide du mot-clé `return`.

Une fonction est invoquée en spécifiant le nom de la fonction et en passant des arguments pour chacun des paramètres de la fonction.

Notez que Go prend en charge deux types de commentaires.
Les commentaires sur une seule ligne sont précédés de `//` et les commentaires sur plusieurs lignes sont insérés entre `/*` et `*/`.

```go
package greeting

// Hello is a public function.
func Hello (name string) string {
    return hi(name)
}

// hi is a private function.
func hi (name string) string {
    return "hi " + name
}
```

## Instructions

Dans cet exercice, vous allez écrire du code pour vous aider à cuisiner de superbes lasagnes à partir de votre livre de cuisine préféré.

Vous avez quatre tâches, toutes liées au temps passé à cuisiner les lasagnes.

## 1. Définissez la durée prévue du four en minutes

Définissez la constante `OvenTime` avec le nombre de minutes pendant lesquelles les lasagnes doivent rester au four. Selon le livre de cuisine, la durée de four prévue en minutes est de 40 :

```go
OvenTime
// => 40
```

## 2. Calculez le temps de four restant en minutes

Définissez la fonction `RemainingOvenTime()` qui prend comme paramètre les minutes réelles pendant lesquelles les lasagnes ont été dans le four et renvoie le nombre de minutes pendant lesquelles les lasagnes doivent encore rester dans le four, en fonction de la durée de four prévue en minutes de la tâche précédente.

```go
func RemainingOvenTime(actual int) int {
    // TODO
}

RemainingOvenTime(30)
// => 10
```

## 3. Calculez le temps de préparation en minutes

Définissez la fonction `PreparationTime` qui prend comme paramètre le nombre de couches que vous avez ajoutées à la lasagne et renvoie le nombre de minutes que vous avez passées à préparer la lasagne, en supposant que chaque couche vous prend 2 minutes à préparer.

```go
func PreparationTime(numberOfLayers int) int {
    // TODO
}

PreparationTime(2)
// => 4
```

## 4. Calculez le temps de travail écoulé en minutes

Définissez la fonction `ElapsedTime` qui prend deux paramètres : le premier paramètre est le nombre de couches que vous avez ajoutées à la lasagne et le deuxième paramètre est le nombre de minutes pendant lesquelles la lasagne a été au four.
La fonction doit indiquer le nombre total de minutes que vous avez travaillées sur la cuisson des lasagnes, qui est la somme du temps de préparation en minutes et du temps en minutes que la lasagne a passé au four à ce moment-là.

```go
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    // TODO
}

ElapsedTime(3, 20)
// => 26
```

## Source

### Créé par

- @tehsphinx

### Contribué à par

- @ekingery
- @andrerfcsantos
- @bobtfish
