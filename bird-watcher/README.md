# Observateur d'Oiseaux

Bienvenue dans « Observateur d'Oiseaux » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Syntaxe générale

La boucle for est l'une des instructions les plus couramment utilisées pour exécuter à plusieurs reprises une logique. En Go, elle se compose du mot-clé `for`, d'un en-tête et d'un bloc de code contenant le corps de la boucle enveloppé dans des accolades. L'en-tête se compose de 3 composants séparés par des points-virgules `;` : init, condition et post.

```go
for init; condition; post {
  // corps de la boucle - code exécuté à plusieurs reprises tant que la condition est vraie
}
```

- Le composant **init** est du code qui s'exécute une seule fois avant le démarrage de la boucle.
- Le composant **condition** doit être une expression qui s'évalue à un booléen et contrôle quand la boucle doit s'arrêter. Le code à l'intérieur de la boucle s'exécutera tant que cette condition s'évalue à vrai. Dès que cette expression s'évalue à faux, aucune autre itération de la boucle ne s'exécutera.
- Le composant **post** est du code qui s'exécutera à la fin de chaque itération.

**Remarque :** Contrairement à d'autres langages, il n'y a pas de parenthèses `()` entourant les trois composants de l'en-tête. En fait, l'insertion de telles parenthèses est une erreur de compilation. Cependant, les accolades `{ }` entourant le corps de la boucle sont toujours requises.

## Boucles For - Un exemple

Le composant init configure généralement une variable compteur, la condition vérifie si la boucle doit continuer ou s'arrêter, et le composant post incrémente généralement le compteur à la fin de chaque répétition.

```go
for i := 1; i < 10; i++ {
  fmt.Println(i)
}
```

Cette boucle affichera les nombres de `1` à `9` (y compris `9`).
Définir l'étape se fait souvent en utilisant une instruction d'incrémentation ou de décrémentation, comme montré dans l'exemple ci-dessus.

## Instructions

Vous êtes un observateur d'oiseaux passionné qui garde une trace du nombre d'oiseaux qui ont visité votre jardin.
Habituellement, vous utilisez un comptage dans un carnet pour compter les oiseaux, mais pour mieux travailler avec vos données, vous avez numérisé les comptages d'oiseaux des semaines passées.

## 1. Déterminer le nombre total d'oiseaux comptés jusqu'à présent

Commençons par analyser les données en obtenant une vue de haut niveau.
Découvrez combien d'oiseaux vous avez comptés au total depuis le début de vos journaux.

Implémentez une fonction `TotalBirdCount` qui accepte une tranche d'`int`s contenant le compte d'oiseaux par jour.
Elle doit retourner le nombre total d'oiseaux que vous avez comptés.

```go
birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
TotalBirdCount(birdsPerDay)
// => 34
```

## 2. Calculer le nombre d'oiseaux visiteurs dans une semaine spécifique

Maintenant que vous avez une compréhension générale de vos chiffres de comptage d'oiseaux, vous voulez faire une analyse plus fine.

Implémentez une fonction `BirdsInWeek` qui accepte une tranche de comptes d'oiseaux par jour et un numéro de semaine.

Elle retourne le nombre total d'oiseaux que vous avez comptés dans cette semaine spécifique.
Vous pouvez supposer que les semaines sont toujours suivies complètement.

```go
birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
BirdsInWeek(birdsPerDay, 2)
// => 12
```

## 3. Corriger une erreur de comptage

Vous avez réalisé que tout le temps où vous essayiez de tenir un registre des oiseaux, il y avait un oiseau qui se cachait dans un coin éloigné du jardin.

Vous avez découvert que cet oiseau a toujours passé tous les deux jours dans votre jardin.

Vous ne savez pas exactement où il était pendant ces jours, mais certainement pas dans votre jardin.

Votre intuition d'observateur d'oiseaux vous dit également que l'oiseau était dans votre jardin le premier jour que vous avez suivi dans votre liste.

Compte tenu de ces nouvelles informations, écrivez une fonction `FixBirdCountLog` qui prend une tranche d'oiseaux comptés par jour comme argument et retourne la tranche après correction de l'erreur de comptage.

```go
birdsPerDay := []int{2, 5, 0, 7, 4, 1}
FixBirdCountLog(birdsPerDay)
// => [3 5 1 7 5 1]
```

## Source

### Créé par

- @sachinmk27

### Contributeurs

- @andrerfcsantos
