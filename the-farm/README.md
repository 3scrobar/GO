# La Ferme

Bienvenue dans « La Ferme » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## L'interface d'erreur

La gestion des erreurs n'est **pas** effectuée via des exceptions dans Go.
Au lieu de cela, les erreurs sont des _valeurs_ normales de types qui implémentent l'interface `error` intégrée.
L'interface `error` est très minimale.
Il ne contient qu'une seule méthode `Error()` qui renvoie le message d'erreur sous forme de chaîne.

```go
type error interface {
  Error() string
}
```

Chaque fois que vous définissez une fonction dans laquelle une erreur pourrait se produire lors de l'exécution et qui doit atteindre l'appelant, vous devez inclure `error` comme l'un des types de retour.
Si la fonction a plusieurs valeurs de retour, par convention `error` est toujours la dernière.

```go
func DoSomething() (int, error) {
  // ...
}
```

## Créer et renvoyer une erreur simple

Vous n'êtes pas obligé de toujours implémenter vous-même l'interface d'erreur.
Pour créer une erreur simple, vous pouvez utiliser la fonction `errors.New()` qui fait partie du package de bibliothèque standard `errors`.
La seule chose que vous devez transmettre est le message d'erreur sous forme de chaîne, et `errors.New()` se chargera de créer une valeur qui contient votre message et implémentera l'interface `error`.

Si la fonction renvoie une erreur, il est recommandé de renvoyer la valeur zéro pour tous les autres paramètres de retour :

```go
func DoSomething() (SomeStruct, int, error) {
  // ...
  return SomeStruct{}, 0, errors.New("failed to calculate result")
}
```

~~~~exercism/caution
You should not assume that all functions return zero values for other return values if an error is present.
It is best practice to assume that it is not safe to use any of the other return values if an error occurred.
The only exceptions are cases where the documentation clearly states that other returns values are meaningful in case of an error.
~~~~

Si vous souhaitez utiliser une erreur aussi simple à plusieurs endroits, vous devez déclarer une variable pour l'erreur au lieu d'utiliser `errors.New` en ligne.
Par convention, le nom de la variable doit commencer par `Err` ou `err` (selon qu'elle est exportée ou non).
Ces variables d'erreur sont souvent appelées _erreurs sentinelles_.

```go
import "errors"

var ErrNotFound = errors.New("resource was not found")

func DoSomething() error {
  // ...
  return ErrNotFound
}
```

Renvoyez `nil` pour l'erreur afin de signaler qu'il n'y a eu aucune erreur lors de l'exécution de la fonction :

```go
func Foo() (int, error) {
  return 10, nil
}
```

## Vérification des erreurs

Si vous appelez une fonction qui renvoie une erreur, il est courant de stocker la valeur de l'erreur dans une variable appelée `err`.
Avant d'utiliser le résultat réel de la fonction, vous devez vérifier qu'il n'y a pas eu d'erreur.

Pour éviter d'imbriquer le « chemin heureux » de votre code, les cas d'erreur doivent être traités en premier.
Nous pouvons utiliser `==` et `!=` pour comparer l'erreur avec `nil` et nous savons qu'il y a eu une erreur lorsque `err` n'est pas `nil`.

```go
func processUserFile() error {
	file, err := os.Open("./users.csv")
	if err != nil {
		return err
	}

	// do something with file
}
```

La plupart du temps, l'erreur sera renvoyée dans la pile de fonctions, comme indiqué dans l'exemple ci-dessus.
Une autre façon de gérer l'erreur pourrait être de la consigner et de poursuivre une autre opération.
Il est recommandé de renvoyer ou de consigner l'erreur, jamais les deux.

Étant donné que la plupart des fonctions de Go incluent une erreur comme l'une des valeurs de retour, vous verrez/utiliserez le modèle `if err != nil` partout dans le code Go.

## Types d'erreurs personnalisés

Si vous souhaitez que votre erreur inclue plus d'informations que la simple chaîne du message d'erreur, vous pouvez créer un type d'erreur personnalisé.
Comme mentionné précédemment, tout ce qui implémente l'interface `error` (c'est-à-dire qui possède une méthode `Error() string`) peut servir d'erreur dans Go.

Habituellement, une structure est utilisée pour créer un type d'erreur personnalisé.
Par convention, les noms de types d'erreur personnalisés doivent se terminer par `Error`.
De plus, il est préférable de configurer la méthode `Error() string` avec un récepteur de pointeur, voir ce [commentaire Stackoverflow][stackoverflow-errors] pour en savoir plus sur le raisonnement.
Notez que cela signifie que vous devez renvoyer un pointeur vers votre erreur personnalisée, sinon elle ne comptera pas comme `error` car la valeur non-pointeur ne fournit pas la méthode `Error() string`.

```go
type MyCustomError struct {
  message string
  details string
}

func (e *MyCustomError) Error() string {
  return fmt.Sprintf("%s, details: %s", e.message, e.details)
}

func someFunction() error {
  // ...
  return &MyCustomError{
    message: "...",
    details: "...",
  }
}
```

[stackoverflow-errors]: https://stackoverflow.com/a/50333850

## Instructions

Le jour que vous avez tant attendu est enfin arrivé et vous êtes désormais l'heureux propriétaire d'une belle ferme dans les Alpes.

Vous n'aimez toujours pas vous réveiller trop tôt le matin pour nourrir vos vaches.
Parce que vous êtes un excellent ingénieur, vous construisez un distributeur de nourriture, le `FEED-M-ALL`.

La dernière chose dont vous avez besoin pour terminer votre projet est un morceau de code qui calcule la quantité de fourrage que chaque vache doit recevoir.
Il est important que chaque vache reçoive le même montant, il faut éviter les conflits.
Les vaches sont très sensibles.

Heureusement, vous n'avez pas besoin d'élaborer vous-même toutes les formules pour calculer les quantités de fourrage.
Vous utilisez une mystérieuse bibliothèque externe que vous avez trouvée sur Internet.
Cela est censé donner les vaches les plus heureuses.
La bibliothèque expose un type qui remplit l'interface suivante.
Vous vous y fierez dans le code que vous écrivez vous-même.

```go
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}
```

Au fur et à mesure que vous travaillez sur votre code, vous améliorerez la gestion des erreurs pour le rendre plus robuste et plus facile à déboguer ultérieurement lorsque vous l'utiliserez dans votre vie quotidienne de ferme.

## 1. Répartissez la nourriture uniformément

Tout d’abord, vous vous concentrez sur l’écriture du code nécessaire pour calculer la quantité de fourrage par vache.

Implémentez une fonction `DivideFood` qui accepte un `FodderCalculator` et un nombre de vaches sous forme d'entier comme arguments.
*Pour cette tâche, vous supposez que le nombre de vaches transmises est toujours supérieur à zéro.*
La fonction doit renvoyer la quantité de nourriture par vache sous la forme d'un `float64` ou d'une erreur le cas échéant.

Pour faire le calcul, il faut d’abord récupérer la quantité totale de fourrage pour toutes les vaches.
Cela se fait en appelant la méthode `FodderAmount` et en transmettant le nombre de vaches.
De plus, vous avez besoin d’un facteur avec lequel ce montant doit être multiplié.
Vous obtenez ce facteur en appelant la méthode `FatteningFactor`.
Avec ces deux valeurs et le nombre de vaches, vous pouvez désormais calculer la quantité de nourriture par vache (sous forme de `float64`).
C'est ce qui doit être renvoyé par la fonction `DivideFood`.

Si l'une des méthodes que vous appelez renvoie une erreur, l'exécution doit s'arrêter et cette erreur doit être renvoyée (telle quelle) par la fonction `DivideFood`.

```go
// For this example, we assume FodderAmount returns 50
// and FatteningFactor returns 1.5.
DivideFood(fodderCalculator, 5)
// => 15 <nil>

// Now assuming FodderAmount returns an error with message "something went wrong".
DivideFood(fodderCalculator, 5)
// => 0 "something went wrong"
```

## 2. Vérifiez le nombre de vaches

En travaillant sur la première tâche ci-dessus, vous avez réalisé que la bibliothèque externe que vous utilisez n'est pas d'aussi haute qualité que vous le pensiez.
Par exemple, il ne peut pas gérer correctement les entrées non valides.
Vous souhaitez contourner cette limitation en ajoutant une vérification de la valeur d'entrée dans votre propre code.

Écrivez une fonction `ValidateInputAndDivideFood` qui a la même signature que `DivideFood` ci-dessus.

- Si le nombre de vaches transmises est supérieur à 0, la fonction doit appeler `DivideFood` et renvoyer les résultats de cet appel.
- Si le nombre de vaches est de 0 ou moins, la fonction doit renvoyer une erreur avec le message `"invalid number of cows"`.

```go
ValidateInputAndDivideFood(fodderCalculator, 5)
// => 15 <nil>

ValidateInputAndDivideFood(fodderCalculator, -2)
// => 0 "invalid number of cows"
```

## 3. Améliorer la gestion des erreurs

Vérifier le nombre de vaches avant de le transmettre était une bonne chose, mais vous n'êtes pas tout à fait satisfait du message d'erreur non spécifique.
Vous décidez de faire mieux en créant un type d'erreur personnalisé appelé `InvalidCowsError`.

L'erreur personnalisée doit contenir le nombre de vaches (`int`) et un message personnalisé (`string`) et la méthode `Error` doit sérialiser les données au format suivant :
```txt
{number of cows} cows are invalid: {custom message}
```

Equipé de votre erreur personnalisée, implémentez une fonction `ValidateNumberOfCows` qui accepte le nombre de vaches comme un entier et renvoie une erreur (ou zéro).

- Si le nombre de vaches est inférieur à 0, la fonction renvoie un `InvalidCowsError` avec le message personnalisé défini sur `"there are no negative cows"`.
- Si le nombre de vaches est 0, la fonction renvoie un `InvalidCowsError` avec le message personnalisé défini sur `"no cows don't need food"`.
- Sinon, la fonction renvoie `nil` pour indiquer que la validation a réussi.

```go
err := ValidateNumberOfCows(-5)
err.Error()
// => "-5 cows are invalid: there are no negative cows"
```

Après le dur travail de mise en place de cette fonction de validation, vous remarquez que c'est déjà le soir et vous quittez votre bureau pour profiter du coucher de soleil sur les montagnes.
Vous laissez la tâche d'ajouter la nouvelle fonction de validation à votre code à un autre jour.

## Source

### Créé par

- @brugnara
- @jmrunkle
- @junedev
