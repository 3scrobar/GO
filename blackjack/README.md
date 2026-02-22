# Blackjack

Bienvenue dans « Blackjack » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Comme d'autres langages, Go fournit également une instruction `switch`. Les instructions Switch sont un moyen plus court d'écrire de longues instructions `if ... else if`. Pour effectuer un changement, nous commençons par utiliser le mot-clé `switch` suivi d'une valeur ou d'une expression. On déclare ensuite chacune des conditions avec le mot clé `case`. Nous pouvons également déclarer un cas `default`, qui s'exécutera lorsqu'aucune des conditions `case` précédentes ne correspond :

```go
operatingSystem := "windows"

switch operatingSystem {
case "windows":
    // do something if the operating system is windows
case "linux":
    // do something if the operating system is linux
case "macos":
    // do something if the operating system is macos
default:
    // do something if the operating system is none of the above
} 
```

Une chose intéressante à propos des instructions switch est que la valeur après le mot-clé `switch` peut être omise et nous pouvons avoir des conditions booléennes pour chaque `case` :

```go
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
```

## Instructions

Dans cet exercice, nous allons simuler le premier tour d'une partie [Blackjack](https://en.wikipedia.org/wiki/Blackjack).

Vous recevrez deux cartes et pourrez voir la carte face visible du croupier. Toutes les cartes sont représentées par une chaîne telle que « as », « roi », « trois », « deux », etc. Les valeurs de chaque carte sont :

| carte | valeur | carte | valeur |
| :---: | :---: | :-----: | :---: |
|  as |  11 | huit |   8 |
|  deux |   2 | neuf |   9 |
| trois |   3 |  dix |  10 |
| quatre |   4 | prise |  10 |
| cinq |   5 | reine |  10 |
|  six |   6 | roi |  10 |
| sept |   7 | *autre* |   0 |

**Remarque** : Généralement, les as peuvent prendre la valeur de 1 ou 11, mais par souci de simplicité, nous supposerons qu'ils ne peuvent prendre que la valeur de 11.

En fonction de vos deux cartes et de la carte du croupier, il existe une stratégie pour le premier tour de jeu, dans laquelle vous disposez des options suivantes :

- Soutien (S)
- Frapper (H)
- Divisé (P)
- Gagner automatiquement (W)

Bien que ce ne soit pas encore optimal, vous suivrez la stratégie développée par votre ami Alex, qui est la suivante :

- Si vous avez une paire d'as, vous devez toujours les partager.
- Si vous avez un Blackjack (deux cartes qui totalisent une valeur de 21) et que le croupier n'a pas d'as, de figure (Valet/Dame/Roi) ou de dix alors vous gagnez automatiquement. Si le croupier possède l'une de ces cartes, vous devrez attendre que l'autre carte soit révélée.
- Si vos cartes totalisent une valeur comprise entre [17, 20], vous devez toujours rester debout.
- Si vos cartes totalisent une valeur comprise entre [12, 16], vous devez toujours rester debout, sauf si le croupier a un 7 ou plus, auquel cas vous devez toujours tirer.
- Si vos cartes totalisent 11 ou moins, vous devriez toujours tirer.

## 1. Calculez la valeur d'une carte donnée.

Implémentez une fonction pour calculer la valeur numérique d'une carte :

```go
value := ParseCard("ace")
fmt.Println(value)
// Output: 11
```

## 2. Implémentez la logique de décision pour le premier tour.

Écrivez une fonction qui implémente la logique de décision comme décrit ci-dessus :

```go
func FirstTurn(card1, card2, dealerCard string) string
```

Voici quelques exemples des résultats attendus :

```go
FirstTurn("ace", "ace", "jack") == "P"
FirstTurn("ace", "king", "ace") == "S"
FirstTurn("five", "queen", "ace") == "H"
```

## Source

### Créé par

- @andres-zartab

### Contribué à par

- @tehsphinx
- @andrerfcsantos
- @norbs57
