# Journaux, Journaux, Journaux !

Bienvenue dans « Journaux, Journaux, Journaux ! » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Le type `rune` en Go est un alias pour `int32`.
Compte tenu de ce type `int32` sous-jacent, le type `rune` contient une valeur entière signée de 32 bits.
Cependant, contrairement à un type `int32`, la valeur entière stockée dans un type `rune` représente un seul caractère Unicode.

## Unicode et points de code Unicode

Unicode est un sur-ensemble d'ASCII qui représente les caractères en attribuant un numéro unique à chaque caractère.
Ce numéro unique s'appelle un point de code Unicode.
Unicode vise à représenter tous les caractères du monde, y compris diverses alphabets, chiffres, symboles et même emoji sous forme de points de code Unicode.

En Go, le type `rune` représente un seul point de code Unicode.

Le tableau suivant contient des exemples de caractères Unicode avec leurs points de code Unicode et valeurs décimales :

| Caractère Unicode | Point de code Unicode | Valeur décimale |
|-------------------|-----------------------|-----------------|
| 0| `U+0030` | `48` |
| Un | `U+0041` | `65` |
| un | `U+0061` | `97` |
| | `U+00BF` | `191` |
| p | `U+03C0` | `960` |
| 🧠 | `U+1F9E0` | `129504` |

##UTF-8

UTF-8 est un encodage de caractères à largeur variable utilisé pour encoder chaque point de code Unicode en 1, 2, 3 ou 4 octets.
Comme un point de code Unicode peut être encodé en maximum 4 octets, le type `rune` doit pouvoir contenir jusqu'à 4 octets de données.
C'est pourquoi le type `rune` est un alias pour `int32` car un type `int32` peut contenir jusqu'à 4 octets de données.

Les fichiers de code source Go sont encodés en UTF-8.

## Utilisation des Runes

Les variables de type `rune` sont déclarées en plaçant un caractère entre guillemets simples :

```go
myRune := '¿'
```

Puisque `rune` n'est qu'un alias pour `int32`, imprimer le type d'une rune donnera `int32` :

```go
myRune := '¿'
fmt.Printf("myRune type: %T\n", myRune)
// Output: myRune type: int32
```

De même, imprimer la valeur d'une rune donnera sa valeur entière (décimale) :

```go
myRune := '¿'
fmt.Printf("myRune value: %v\n", myRune)
// Output: myRune value: 191
```

Pour imprimer le caractère Unicode représenté par la rune, utilisez le verbe de formatage `%c` :

```go
myRune := '¿'
fmt.Printf("myRune Unicode character: %c\n", myRune)
// Output: myRune Unicode character: ¿
```

Pour imprimer le point de code Unicode représenté par la rune, utilisez le verbe de formatage `%U` :

```go
myRune := '¿'
fmt.Printf("myRune Unicode code point: %U\n", myRune)
// Output: myRune Unicode code point: U+00BF
```

## Runes et Chaînes

Les chaînes en Go sont encodées en UTF-8, ce qui signifie qu'elles contiennent des caractères Unicode.
Les caractères dans les chaînes sont stockés et encodés en 1, 2, 3 ou 4 octets selon le caractère Unicode qu'ils représentent.

En Go, les tranches sont utilisées pour représenter des séquences et ces tranches peuvent être itérées à l'aide de range.
Lorsque nous itérons sur une chaîne, Go convertit la chaîne en une série de Runes, dont chacune fait 4 octets (rappelez-vous, le type rune est un alias pour un `int32` !)

Bien qu'une chaîne ne soit qu'une tranche d'octets, le mot-clé `range` itère sur les runes d'une chaîne, pas ses octets.

Dans cet exemple, la variable `index` représente l'index de départ de la séquence d'octets de la rune actuelle et la variable `char` représente la rune actuelle :

```go
myString := "❗hello"
for index, char := range myString {
  fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\n", index, char, char)
}
// Output:
// Index: 0	Character: ❗		Code Point: U+2757
// Index: 3	Character: h		Code Point: U+0068
// Index: 4	Character: e		Code Point: U+0065
// Index: 5	Character: l		Code Point: U+006C
// Index: 6	Character: l		Code Point: U+006C
// Index: 7	Character: o		Code Point: U+006F
```

Puisque les runes peuvent être stockées en 1, 2, 3 ou 4 octets, la longueur d'une chaîne peut ne pas toujours égaler le nombre de caractères dans la chaîne.
Utilisez la fonction intégrée `len` pour obtenir la longueur d'une chaîne en octets et la fonction `utf8.RuneCountInString` pour obtenir le nombre de runes dans une chaîne :

```go
import "unicode/utf8"

myString := "❗hello"
stringLength := len(myString)
numberOfRunes := utf8.RuneCountInString(myString)

fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)
// Output: myString - Length: 8 - Runes: 6
```

## Instructions

Vous avez été chargé de créer une bibliothèque de journaux pour aider à gérer les journaux de votre organisation. Cette bibliothèque permettra aux utilisateurs d'identifier quelle application a émis un journal donné, de corriger les journaux corrompus et de déterminer si une ligne de journal donnée se trouve dans une certaine limite de caractères.

## 1. Identifier quelle application a émis un journal

Les journaux proviennent de plusieurs applications qui utilisent chacune son propre format de journal propriétaire. L'application qui émet un journal doit être identifiée avant qu'elle ne soit stockée dans un système d'agrégation de journaux.

Implémentez la fonction `Application` qui prend une ligne de journal et retourne l'application qui a émis la ligne de journal.

Pour identifier quelle application a émis une ligne de journal donnée, recherchez dans la ligne de journal un caractère spécifique selon le tableau suivant :

| Application      | Caractère | Point de code Unicode |
|------------------|-----------|----------------------|
| `recommendation` | ❗ | `U+2757` |
| `search` | 🔍 | `U+1F50D` |
| `weather` | ☀ | `U+2600` |

Si une ligne de journal ne contient pas l'un des caractères du tableau ci-dessus, retournez `default` à l'appelant. Si une ligne de journal contient plus d'un caractère du tableau ci-dessus, retournez l'application correspondant au premier caractère trouvé dans la ligne de journal de gauche à droite.

```go
Application("❗ recommended search product 🔍")
// => recommendation
```

## 2. Corriger les journaux corrompus

En raison d'un bug rare mais persistant dans l'infrastructure de journalisation, certains caractères dans les journaux peuvent être corrompus. Après avoir passé du temps à identifier les caractères corrompus et leur valeur d'origine, vous décidez de mettre à jour la bibliothèque de journaux pour aider à corriger les journaux corrompus.

Implémentez la fonction `Replace` qui prend une ligne de journal, un caractère corrompu, et la valeur d'origine et retourne une ligne de journal modifiée qui a tous les occurrences du caractère corrompu remplacés par la valeur d'origine.

```go
log := "please replace '👎' with '👍'"

Replace(log, '👎', '👍')
// => please replace '👍' with '👍'"
```

## 3. Déterminer si un journal peut être affiché

Les systèmes responsables de l'affichage des journaux ont une limite sur le nombre de caractères pouvant être affichés par ligne de journal. Ainsi, les utilisateurs demandent que cette bibliothèque inclue une fonction d'assistance pour déterminer si une ligne de journal se trouve dans une limite de caractères spécifique.

Implémentez la fonction `WithinLimit` qui prend une ligne de journal et une limite de caractères et retourne si oui ou non la ligne de journal se trouve dans la limite de caractères.

```go
WithinLimit("hello❗", 6)
// => true
```

## Source

### Créé par

- @sudomateo
- @tehsphinx
