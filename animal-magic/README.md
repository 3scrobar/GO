# Animal Magic

Bienvenue sur Animal Magic sur la piste Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans l'utiliser :)

## Introduction

Le paquet [math/rand][mathrand] fournit un support pour générer des nombres pseudoaléatoires.

Voici comment générer un entier aléatoire entre `0` et `99` :

```go
n := rand.Intn(100) // n est un int aléatoire, 0 <= n < 100
```

La fonction `rand.Float64` retourne un nombre à virgule flottante aléatoire entre `0.0` et `1.0` :

```go
f := rand.Float64() // f est un float64 aléatoire, 0.0 <= f < 1.0
```

Il y a aussi un support pour mélanger une tranche (ou d'autres structures de données) :

```go
x := []string{"a", "b", "c", "d", "e"}
// mélanger la tranche place ses éléments dans un ordre aléatoire
rand.Shuffle(len(x), func(i, j int) {
	x[i], x[j] = x[j], x[i]
})
```

## Graines

Les séquences de nombres générées par le paquet `math/rand` ne sont pas vraiment aléatoires.
Avec une valeur de "graine" spécifique, les résultats sont entièrement déterministes.

En Go 1.20+, la graine est automatiquement choisie au hasard, vous verrez donc une séquence différente de nombres aléatoires à chaque exécution de votre programme.

Dans les versions antérieures de Go, la graine était `1` par défaut.
Pour obtenir des séquences différentes lors de diverses exécutions du programme, vous deviez initialiser manuellement le générateur de nombres aléatoires, par exemple avec l'heure actuelle, avant de récupérer des nombres aléatoires.

```go
rand.Seed(time.Now().UnixNano())
```

[mathrand]: https://pkg.go.dev/math/rand

## Instructions

Elaine travaille sur un nouveau jeu pour enfants mettant en vedette des animaux et des baguettes magiques.
Il est temps de coder des fonctions pour lancer un dé, générer de l'énergie de baguette aléatoire et mélanger une tranche.

## 1. Lancer un dé.

Implémentez une fonction `RollADie`.

Ce sera le traditionnel dé à vingt faces avec les numéros 1 à 20.

```go
d := RollADie() // d se verra attribuer un int aléatoire, 1 <= d <= 20
```

## 2. Générer de l'énergie de baguette.

Implémentez une fonction `GenerateWandEnergy`.
L'énergie de la baguette doit être un nombre à virgule flottante aléatoire supérieur ou égal à 0.0 et inférieur à 12.0.

```go
f := GenerateWandEnergy()  // f se verra attribuer un float64 aléatoire, 0.0 <= f < 12.0
```

## 3. Mélanger une tranche.

Le jeu comporte huit animaux différents :

- fourmi
- castor
- chat
- chien
- éléphant
- renard
- girafe
- hérisson

Écrivez une fonction `ShuffleAnimals` qui retourne une tranche avec les huit animaux dans un ordre aléatoire.

## Source

### Créé par

- @norbs57

### Contribué par

- @junedev