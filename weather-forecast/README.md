# Prévisions Météorologiques

Bienvenue dans « Prévisions Météorologiques » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Dans l'exercice précédent, nous avons vu qu'il existe deux façons d'écrire des commentaires en Go : les commentaires sur une seule ligne précédés par `//`, et les blocs de commentaires multilignés enveloppés avec `/*` et `*/`.

## Commentaires de documentation

En Go, les commentaires jouent un rôle important dans la documentation du code. Ils sont utilisés par la commande `godoc`, qui extrait ces commentaires pour créer de la documentation sur les paquets Go. Un commentaire de documentation doit être une phrase complète qui commence par le nom de la chose décrite et se termine par un point.

Les commentaires doivent précéder les paquets ainsi que les identifiants exportés, par exemple les fonctions exportées, les méthodes, les variables de paquet, les constantes et les structs, que vous en apprendrez plus dans les exercices suivants.

Une variable au niveau du paquet peut ressembler à ceci :

```go
// TemperatureCelsius représente une certaine température en degrés Celsius.
var TemperatureCelsius float64
```

## Commentaires de paquet

Les commentaires de paquet doivent être écrits directement avant une clause de paquet (`package x`) et commencer par `Package x ...` comme ceci :

```go
// Package kelvin fournit des outils pour convertir
// les températures vers et depuis Kelvin.
package kelvin
```

## Commentaires de fonction

Un commentaire de fonction doit être écrit directement avant la déclaration de fonction. C'est une phrase complète qui commence par le nom de la fonction. Par exemple, un commentaire exporté pour la fonction `Calculate` devrait prendre la forme `Calculate ...`. Il devrait également expliquer quels arguments la fonction prend, ce qu'elle en fait, et ce que ses valeurs de retour signifient, se terminant par un point) :

```go
// CelsiusFreezingTemp retourne une valeur entière égale à la température à laquelle l'eau gèle en degrés Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
```

## Instructions

Goblinocus est un pays qui prend ses prévisions météorologiques très au sérieux. Puisque vous êtes un développeur réputé, responsable et compétent, ils vous ont demandé d'écrire un programme capable de prévoir les conditions météorologiques actuelles de diverses villes de Goblinocus. Vous étiez occupé à l'époque et avez demandé à l'un de vos amis de faire le travail à la place. Après un moment, le président de Goblinocus vous a contacté et a dit qu'il ne comprenait pas le code de votre ami. Quand vous vérifiez le code, vous découvrez que votre ami n'a pas agi comme un programmeur responsable et il n'y a pas de commentaires dans le code. Vous vous sentez obligé de clarifier le programme pour que les gobelins le comprennent aussi.

## 1. Documenter le paquet weather

Puisque les gobelins ne sont pas aussi intelligents que vous, ils ont oublié à quoi devrait servir le paquet. Veuillez écrire un commentaire pour `package weather` qui décrit son contenu. Le commentaire de paquet doit présenter le paquet et fournir des informations pertinentes au paquet dans son ensemble.

## 2. Documenter les variables CurrentCondition et CurrentLocation

Le président de Goblinocus est un peu paranoïaque et craint que les variables sans commentaires ne soient utilisées pour détruire son pays. Veuillez clarifier l'utilisation des variables de paquet `CurrentCondition` et `CurrentLocation` et rassurez l'esprit du président. Cela devrait dire à tout utilisateur du paquet quelles informations les variables stockent et ce qu'il peut en faire.

## 3. Documenter la fonction Forecast()

Les opérateurs de prévisions de Goblinocus veulent savoir ce que la fonction `Forecast()` fait (mais ne leur dites pas comment elle fonctionne, car malheureusement, ils seront encore plus confus). Veuillez écrire un commentaire pour cette fonction qui décrit ce que la fonction fait, mais pas comment elle le fait.

## Source

### Créé par

- @nikimanoledaki
- @micuffaro
