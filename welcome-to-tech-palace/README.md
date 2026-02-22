# Bienvenue au Tech Palace !

Bienvenue dans « Bienvenue au Tech Palace ! » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Un `string` dans Go est une séquence d'octets immuable, qui ne doit pas nécessairement représenter des caractères.

Une chaîne littérale est définie entre guillemets :

```go
const name = "Jane"
```

Les chaînes peuvent être concaténées via l'opérateur `+` :

```go
"Jane" + " " + "Austen"
// => "Jane Austen"
```

Certains caractères spéciaux doivent être échappés avec une barre oblique inverse, tels que `\t` pour un onglet et `\n` pour une nouvelle ligne dans les chaînes.

```go
"How is the weather today?\nIt's sunny"  
// =>
// How is the weather today?
// It's sunny
```

Le package `strings` contient de nombreuses fonctions utiles pour travailler sur les chaînes.
Pour plus d'informations sur les fonctions de chaîne, consultez le [strings package documentation](https://pkg.go.dev/strings).
Voici quelques exemples :

```go
import "strings"

// strings.ToLower returns the string given as argument with all its characters lowercased
strings.ToLower("MaKEmeLoweRCase")
// => "makemelowercase"

// strings.Repeat returns a string with a substring given as argument repeated many times
strings.Repeat("Go", 3)
// => "GoGoGo"
```

## Instructions

Il y a un magasin d'électroménager appelé "Tech Palace" à proximité.
Le propriétaire du magasin a récemment installé un grand écran pour diffuser des messages marketing et afficher un message d'accueil spécial lorsque les clients scannent leur carte de fidélité à l'entrée.
L'écran se compose de nombreuses petites lumières LED et peut afficher plusieurs lignes de texte.

Le propriétaire du magasin a besoin de votre aide avec le code utilisé pour générer le texte du nouvel affichage.

## 1. Créez le message de bienvenue

Pour la plupart des clients qui scannent leur carte de fidélité, le propriétaire du magasin souhaite voir `Welcome to the Tech Palace, ` suivi du nom du client en majuscules sur l'écran.

Implémentez la fonction `WelcomeMessage` qui accepte le nom du client comme argument `string` et renvoie le message souhaité sous forme de `string`.

```go
WelcomeMessage("Judy")
// => Welcome to the Tech Palace, JUDY
```

## 2. Ajoutez une bordure fantaisie

Pour les clients fidèles qui achètent beaucoup au magasin, le propriétaire souhaite que l'affichage de bienvenue soit plus sophistiqué en ajoutant une ligne d'étoiles avant et après le message de bienvenue.
Ils ne savent pas encore combien d'étoiles doivent figurer dans les lignes, ils veulent donc que cela soit configurable.

Écrivez une fonction `AddBorder` qui accepte un message de bienvenue (un `string`) et le nombre d'étoiles par ligne (tapez `int`) comme arguments.
Il doit renvoyer un `string` composé de 3 lignes, une ligne avec le nombre d'étoiles souhaité, puis le message de bienvenue tel qu'il a été transmis, puis une autre ligne d'étoiles.

```go
AddBorder("Welcome!", 10)
```

Devrait renvoyer ce qui suit :

```
**********
Welcome!
**********
```

## 3. Nettoyer les anciens messages marketing

Avant d'installer ce nouvel écran, le magasin disposait d'un écran similaire qui ne pouvait afficher que des lignes statiques non configurables.
Le propriétaire souhaite réutiliser certains des anciens messages marketing sur le nouvel écran.
Cependant, les données incluent déjà une bordure en étoile et quelques espaces malheureux.
Votre tâche consiste à nettoyer les messages afin qu'ils puissent être réutilisés.

Implémentez une fonction `CleanUpMessage` qui accepte l'ancien message marketing sous forme de chaîne.
La fonction doit d'abord supprimer toutes les étoiles du texte, puis supprimer les espaces de début et de fin du texte restant.
La fonction devrait alors renvoyer le message nettoyé.

```go
message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

CleanUpMessage(message)
// => BUY NOW, SAVE 10%
```

## Source

### Créé par

- @erikschierboom
- @junedev

### Contribué à par

- @kekimaker
