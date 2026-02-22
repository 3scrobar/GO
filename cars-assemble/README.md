# Les voitures s'assemblent

Bienvenue dans « Les voitures s'assemblent » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Numéros

Go contient des types numériques de base qui peuvent représenter des ensembles de valeurs entières ou à virgule flottante.
Il existe différents types en fonction de la taille de valeur dont vous avez besoin et de l'architecture de l'ordinateur sur lequel l'application est exécutée (par exemple 32 bits et 64 bits).

Dans le cadre de cet exercice, vous traiterez uniquement de :

- `int` : par ex. `0`, `255`, `2147483647`. Entier signé d'une taille d'au moins 32 bits (plage de valeurs : -2147483648 à 2147483647).
Mais cela dépendra de l'architecture du système.
La plupart des ordinateurs modernes sont en 64 bits, donc `int` aura une taille de 64 bits (taux de valeur : -9223372036854775808 à 9223372036854775807).

- `float64` : par ex. `0.0`, `3.14`. Contient l’ensemble de tous les nombres à virgule flottante de 64 bits.

- `uint` : par ex. `0`, `255`. Entier non signé de même taille que `int` (plage de valeurs : 0 à 4294967295 pour 32 bits et 0 à 18446744073709551615 pour 64 bits)

Les nombres peuvent être convertis en d'autres types numériques via la conversion de type.

## Opérateurs arithmétiques

Go prend en charge de nombreux opérateurs arithmétiques standards :

| Opérateur | Exemple |
| -------- | ---------- |
| `+` | `4 + 6 == 10` |
| `-` | `15 - 10 == 5` |
| `*` | `2 * 3 == 6` |
| `/` | `13 / 3 == 4` |
| `%` | `13 % 3 == 1` |

Pour une division entière, le reste est supprimé (par exemple `5 / 2 == 2`).

Go a une affectation abrégée pour les opérateurs ci-dessus (par exemple, `a += 5` est l'abréviation de `a = a + 5`).
Go prend également en charge les instructions d'incrémentation et de décrémentation `++` et `--` (par exemple `a++`).

## Conversion entre les types

La conversion entre les types se fait via une fonction avec le nom du type vers lequel convertir.
Par exemple:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (ie. 42.0)
var y float64 = 11.9 // y has type float64
i := int(y) // i has type int (ie. 11)
```
## Opérations arithmétiques sur différents types

Dans de nombreux langages, vous pouvez effectuer des opérations arithmétiques sur différents types de variables, mais en Go, cela génère une erreur.
Par exemple:

```go
var x int = 42

// this line produces an error
value := float32(2.0) * x // invalid operation: mismatched types float32 and int

// you must convert int type to float32 before performing arithmetic operation
value := float32(2.0) * float32(x)
```

## Instructions

Dans cet exercice, vous écrirez du code pour analyser la production dans une usine automobile.

## 1. Calculez le nombre de voitures en état de marche produites par heure

Les voitures sont produites sur une chaîne de montage.
La chaîne de montage a une certaine vitesse, qui peut être modifiée.
Plus la vitesse de la chaîne de montage est rapide, plus de voitures sont produites.
Cependant, la modification de la vitesse de la chaîne de montage modifie également le nombre de voitures produites avec succès, c'est-à-dire des voitures sans aucune erreur dans leur production.

Implémentez une fonction qui prend en compte le nombre de voitures produites par heure et le taux de réussite et calcule le nombre de voitures réussies fabriquées par heure. Le taux de réussite est donné en pourcentage, de `0` à `100` :

```go
CalculateWorkingCarsPerHour(1547, 90)
// => 1392.3
```

**Remarque :** la valeur de retour doit être un `float64`.

## 2. Calculez le nombre de voitures en état de marche produites par minute

Implémentez une fonction qui prend en compte le nombre de voitures produites par heure et le taux de réussite et calcule le nombre de voitures produites avec succès chaque minute :

```go
CalculateWorkingCarsPerMinute(1105, 90)
// => 16
```

**Remarque :** la valeur de retour doit être un `int`.

## 3. Calculer le coût de production

Chaque voiture coûte normalement 10 000 $ à produire individuellement, qu'elle soit réussie ou non.
Mais avec un peu de planification, 10 voitures peuvent être produites ensemble pour 95 000 $.

Par exemple, 37 voitures peuvent être produites de la manière suivante :
37 = 3 x groupes de dix + 7 voitures individuelles

Le coût pour 37 voitures est donc :
3\*95 000+7\*10 000=355 000

Implémentez la fonction `CalculateCost` qui calcule le coût de production d'un certain nombre de voitures, qu'elles réussissent ou non :

```go
CalculateCost(37)
// => 355000

CalculateCost(21)
// => 200000
```

**Remarque :** la valeur de retour doit être un `uint`.

## Source

### Créé par

- @DavyJ0nes
- @kahgoh

### Contribué à par

- @tehsphinx
- @andrerfcsantos
