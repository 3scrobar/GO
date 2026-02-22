# Achat de véhicule

Bienvenue dans « Achat de véhicule » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Comparaison

Les nombres In Go peuvent être comparés à l’aide des opérateurs relationnels et d’égalité suivants.

| Comparaison | Opérateur |
| -------------------| --------- |
| égal | `==` |
| pas égal | `!=` |
| moins | `<` |
| inférieur ou égal | `<=` |
| plus grand | `>` |
| supérieur ou égal | `>=` |

Le résultat de la comparaison est toujours une valeur booléenne, donc `true` ou `false`.

```go
a := 3

a != 4 // true
a > 5  // false
```

Les opérateurs de comparaison ci-dessus peuvent également être utilisés pour comparer des chaînes.
Dans ce cas, un ordre lexicographique (dictionnaire) est appliqué.
Par exemple:

```Go
	"apple" < "banana"  // true
	"apple" > "banana"  // false
```

## Si les déclarations

Les conditions dans Go sont similaires aux conditions dans d’autres langues.
Le type sous-jacent de toute opération conditionnelle est le type `bool`, qui peut avoir la valeur `true` ou `false`.
Les conditions sont souvent utilisées comme mécanismes de contrôle de flux pour vérifier diverses conditions.

Pour vérifier un cas particulier, une instruction `if` peut être utilisée, qui exécute son code si la condition sous-jacente est `true` comme ceci :

```go
var value string

if value == "val" {
    return "was val"
}
```

Dans les scénarios impliquant plusieurs cas, de nombreuses instructions `if` peuvent être enchaînées à l'aide des instructions `else if` et `else`.

```go
var number int
result := "This number is "

if number > 0 {
    result += "positive"
} else if number < 0 {
    result += "negative"
} else {
    result += "zero"
}
```

Les instructions if peuvent également inclure une courte instruction d'initialisation qui peut être utilisée pour initialiser une ou plusieurs variables pour l'instruction if.
Par exemple:

```go
num := 7
if v := 2 * num; v > 10 {
    fmt.Println(v)
} else {
    fmt.Println(num)
}
// Output: 14
```

> Remarque : toutes les variables créées dans l'instruction d'initialisation sont hors de portée après la fin de l'instruction if.

## Instructions

Dans cet exercice, vous allez écrire du code pour vous aider à préparer l’achat d’un véhicule.

Vous avez trois tâches, une pour déterminer si vous avez besoin d'un permis, une pour vous aider à choisir entre deux véhicules et une pour estimer le prix acceptable pour un véhicule d'occasion.

## 1. Déterminez si vous aurez besoin d'un permis de conduire

Certains types de véhicules nécessitent un permis de conduire pour pouvoir les conduire.
Supposons que seuls les types `"car"` et `"truck"` nécessitent une licence, tout le reste peut être utilisé sans licence.

Implémentez la fonction `NeedsLicense(kind)` qui prend le type de véhicule et renvoie un booléen indiquant si vous avez besoin d'un permis pour ce type de véhicule.

```go
needLicense := NeedsLicense("car")
// => true

needLicense = NeedsLicense("bike")
// => false

needLicense = NeedsLicense("truck")
// => true
```

## 2. Choisissez entre deux véhicules potentiels à acheter

Vous évaluez vos options de véhicules disponibles.
Vous parvenez à le réduire à deux options, mais vous avez besoin d’aide pour prendre la décision finale.
Pour cela, implémentez la fonction `ChooseVehicle(option1, option2)` qui prend deux véhicules comme arguments et renvoie une décision qui inclut l'option qui arrive en premier dans l'ordre lexicographique.

```go
vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
// => "Toyota Corolla is clearly the better choice."

ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
// => "Volkswagen Beetle is clearly the better choice."
```

## 3. Calculer une estimation du prix d'un véhicule d'occasion

Maintenant que vous avez pris votre décision, vous voulez vous assurer d’obtenir un prix équitable chez le concessionnaire.
Puisque vous souhaitez acheter un véhicule d’occasion, le prix dépend de l’âge du véhicule.
Pour une estimation approximative, supposons que si le véhicule a moins de 3 ans, il coûte 80 % du prix d'origine qu'il avait lorsqu'il était neuf.
S'il a au moins 10 ans, cela coûte 50 %.
Si le véhicule a au moins 3 ans mais moins de 10 ans, il coûte 70 % du prix d'origine.

Implémentez la fonction `CalculateResellPrice(originalPrice, age)` qui applique cette logique en utilisant `if`, `else if` et `else` (il existe d'autres moyens si vous souhaitez vous entraîner).
Il prend comme arguments le prix d'origine et l'âge du véhicule et renvoie le prix estimé en concession.

```go
CalculateResellPrice(1000, 1)
// => 800

CalculateResellPrice(1000, 5)
// => 700

CalculateResellPrice(1000, 15)
// => 500
```

**Remarque** la valeur de retour est un `float64`.

## Source

### Créé par

- @kahgoh
