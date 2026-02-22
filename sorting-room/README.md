# Salle de Tri

Bienvenue dans « Salle de Tri » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Conversion de types

Go nécessite une conversion explicite entre différents types.
La conversion entre types (aussi appelée **transtypage**) se fait via une fonction portant le nom du type vers lequel convertir.
Par exemple, pour convertir un `int` en `float64`, vous devriez faire ceci :

```go
var x int = 42 // x a le type int
f := float64(x) // f a le type float64 (c.-à-d. 42.0)
```

## Conversion entre types primitifs et chaînes de caractères

Il existe un package `strconv` pour convertir entre les types primitifs (comme `int`) et `string`.

```go
import "strconv"

var intString string = "42"
var i, err = strconv.Atoi(intString)

var number int = 12
var s string = strconv.Itoa(number)
```

## Assertions de type

Les interfaces en Go peuvent introduire une ambiguïté sur le type sous-jacent.
Une assertion de type permet d'extraire la valeur concrète sous-jacente d'une interface en utilisant cette syntaxe : `interfaceVariable.(concreteType)`.

Par exemple :

```go
var input any = 12
number := input.(int)
```

REMARQUE : cela causera une panique si la variable interface ne contient pas une valeur du type concret.

On peut tester si une valeur interface contient un type concret spécifique en utilisant les deux valeurs de retour de l'assertion de type : la valeur sous-jacente et une valeur booléenne qui indique si l'assertion a réussi.
Par exemple :

```go
str, ok := input.(string) // pas de panique si input n'est pas une string
```

Si `input` contient une `string`, alors `str` sera la valeur sous-jacente et `ok` sera vrai.
Si `input` ne contient pas une `string`, alors `str` sera la valeur zéro du type `string` (c.-à-d. `""` - la chaîne vide) et `ok` sera faux.
Aucune panique ne se produit dans tous les cas.

## Commutateurs de Type

Un **commutateur de type** peut effectuer plusieurs assertions de type en série.
Il a la même syntaxe qu'une assertion de type (`interfaceVariable.(concreteType)`), mais le `concreteType` est remplacé par le mot-clé `type`.
Voici un exemple :

```go
var i any = 12 // essayez : 12.3, true, int64(12), []int{}, map[string]int{}

switch v := i.(type) {
case int:
    fmt.Printf("l'entier %d\n", v)
case string:
    fmt.Printf("la chaîne %s\n", v)
default:
    fmt.Printf("type, %T, non traité explicitement : %#v\n", v, v)
}
```

## Instructions

Jen travaille dans la salle de tri d'une grande usine.
La salle de tri doit traiter tout ce qui y arrive en le catégorisant avec une étiquette.
Jen est responsable des choses qui ont été pré-catégorisées comme des nombres et a besoin d'un programme pour l'aider au tri.

La plupart des valeurs primitives doivent obtenir des étiquettes simples.
Pour les nombres, elle veut des chaînes disant `"This is the number 2.0"` (si le nombre était 2).
Jen veut le même résultat pour les entiers et les flottants.

Il existe quelques interfaces `Box` qui doivent être dépliées pour obtenir leur contenu.
Pour une `NumberBox`, elle veut des chaînes disant `"This is a box containing the number 3.0"` (si la méthode `Number()` retourne 3).
Pour une `FancyNumberBox`, elle veut des chaînes disant `"This is a fancy box containing the number 4.0"`, mais seulement si le type est un `FancyNumber`.

Tout ce qui est inattendu doit dire `"Return to sender"` pour que Jen puisse les renvoyer d'où elles viennent.

## 1. Décrire les nombres simples

Jen veut que les nombres retournent des chaînes comme `"This is the number 2.0"` (incluant un chiffre après la décimale) :

```go
fmt.Println(DescribeNumber(-12.345))
// Résultat : This is the number -12.3
```

## 2. Décrire une boîte de nombres

Jen veut que les boîtes de nombres retournent des chaînes comme `"This is a box containing the number 2.0"` (encore une fois, incluant un chiffre après la décimale) :

```go
fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
// Résultat : This is a box containing the number 12.0
```

## 3. Implémenter une méthode d'extraction du nombre d'une boîte de nombres fantaisie

Jen a besoin d'une fonction d'aide pour extraire le nombre d'une `FancyNumberBox`.
Si la `FancyNumberBox` est un `FancyNumber`, extrayez sa valeur et convertissez-la d'une `string` en `int`.
Tout autre type de `FancyNumberBox` doit retourner 0.

```go
fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
// Résultat : 10
fmt.Println(ExtractFancyNumber(AnotherFancyNumber{"4"}))
// Résultat : 0
```

## 4. Décrire une boîte de nombres fantaisie

Si la `FancyNumberBox` est un `FancyNumber`, Jen veut des chaînes disant `"This is a fancy box containing the number 4.0"`.
Tout autre type de `FancyNumberBox` doit dire `"This is a fancy box containing the number 0.0"`.

```go
fmt.Println(DescribeFancyNumberBox(FancyNumber{"10"}))
// Résultat : This is a fancy box containing the number 10.0
fmt.Println(DescribeFancyNumberBox(AnotherFancyNumber{"4"}))
// Résultat : This is a fancy box containing the number 0.0
```

REMARQUE : nous devrions utiliser la fonction `ExtractFancyNumber` !

## 5. Implémenter `DescribeAnything` qui les utilise tous

C'est la fonction principale dont Jen a besoin qui prend n'importe quelle entrée.
`DescribeAnything` doit déléguer aux autres fonctions en fonction du type de la valeur transmise.
Plus précisément :

- `int` et `float64` doivent tous deux déléguer à `DescribeNumber`
- `NumberBox` doit déléguer à `DescribeNumberBox`
- `FancyNumberBox` doit déléguer à `DescribeFancyNumberBox`
- tout le reste doit résulter en `"Return to sender"`

```go
fmt.Println(DescribeAnything(numberBoxContaining{12.345}))
// Résultat : This is a box containing the number 12.3
fmt.Println(DescribeAnything("some string"))
// Résultat : Return to sender
```

## Source

### Créé par

- @jmrunkle

### Contribué par

- @junedev
