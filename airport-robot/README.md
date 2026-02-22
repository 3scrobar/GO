# Robot Aéroportuaire

Bienvenue dans « Robot Aéroportuaire » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Interface comme ensemble de méthodes

Dans sa forme la plus simple, un **type interface** est un ensemble de signatures de méthodes.
Voici un exemple de définition d'interface qui inclut deux méthodes `Add` et `Value` :

```go
type Counter interface {
    Add(increment int)
    Value() int
}
```

Les noms de paramètres comme `increment` peuvent être omis de la définition d'interface, mais ils augmentent souvent la lisibilité.

Les noms d'interface en Go ne contiennent pas le mot `Interface` ou `I`.
Au lieu de cela, ils se terminent souvent par `er`, par exemple `Reader`, `Stringer`.

## Implémenter une interface

Tout type qui définit les méthodes de l'interface implémente automatiquement implicitement l'interface.
Il n'y a pas de mot-clé `implements` en Go.

Le type suivant implémente l'interface `Counter` que nous avons vue ci-dessus.

```go
type Stats struct {
    value int
    // ...
}

func (s Stats) Add(v int) {
    s.value += v
}

func (s Stats) Value() int {
    return s.value
}

func (s Stats) SomeOtherMethod() {
    // Le type peut avoir des méthodes supplémentaires non mentionnées dans l'interface.
}
```

Pour implémenter l'interface, peu importe que la méthode ait un récepteur de valeur ou de pointeur.
(Revisitez le concept [méthodes][concept-methods] si vous en êtes incertain.)

> Une valeur de type interface peut contenir toute valeur qui implémente ces méthodes. [^1]

Cela signifie que `Stats` peut désormais être utilisé dans tous les endroits qui attendent l'interface `Counter`.

```go
func SetUpAnalytics(counter Counter) {
    // ...
}

stats := Stats{}
SetUpAnalytics(stats)
// fonctionne car Stats implémente Counter
```

Parce que les interfaces sont implémentées implicitement, un type peut facilement implémenter plusieurs interfaces.
Il doit seulement avoir toutes les méthodes nécessaires définies.

## Vidéo de l'interface

Il y a une interface très spéciale en Go, l'**interface vide** qui contient zéro méthode.
L'interface vide s'écrit comme ceci : `interface{}`.
Depuis Go 1.18, `any` est un alias pour `interface{}`, et est plus couramment utilisé.

Puisque l'interface vide n'a pas de méthodes, chaque type l'implémente implicitement.
Ceci est utile pour définir une fonction qui peut accepter génériquement toute valeur.
Dans ce cas, le paramètre de fonction utilise le type d'interface vide.

[concepts-méthodes]  : /tracks/go/concepts/methods

## Instructions

Le nouvel aéroport de Berlin a embauché des développeurs pour leur laboratoire de robots et vous commencez votre emploi là-bas.
Ils ont des robots maladroits, un peu ressemblants à des humanoïdes, qu'ils essaient d'utiliser pour améliorer le service à la clientèle.

Votre première tâche au travail est d'écrire un programme permettant au robot de saluer les gens dans leur langue maternelle après avoir scanné leurs passeports au comptoir d'enregistrement automatique.

Le robot est fier de ses capacités, c'est pourquoi il dira toujours d'abord quelle langue il peut parler, puis il saluera la personne.
Par exemple, si quelqu'un scanne un passeport allemand, le robot dirait :

```txt
I can speak German: Hallo Dietrich!
```

## 1. Créer la fonctionnalité abstraite de salutation

Vous n'écrirez pas le code pour les différentes langues vous-même, vous devez donc structurer votre code pour le robot afin que d'autres développeurs puissent facilement ajouter d'autres langues plus tard.

Comme première étape, définissez une interface `Greeter` avec deux méthodes.

- `LanguageName` qui retourne le nom de la langue (une `string`) dans laquelle le robot est censé saluer le visiteur.
- `Greet` qui accepte le nom du visiteur (une `string`) et retourne une `string` avec le message de salutation dans une langue spécifique.

Ensuite, implémentez une fonction `SayHello` qui accepte le nom du visiteur et tout ce qui implémente l'interface `Greeter` comme arguments et retourne la chaîne de salutation désirée.
Par exemple, imaginez une implémentation `Greeter` allemande pour laquelle `LanguageName` retourne `"German"` et `Greet` retourne `"Hallo {name}!"` :

```go
SayHello("Dietrich", germanGreeter)
// => "I can speak German: Hallo Dietrich!"
```

## 2. Implémenter l'italien

Maintenant, votre travail est de faire fonctionner le robot pour les personnes qui scannent des passeports italiens.

Pour cela, créez une struct `Italian` et implémentez les deux méthodes nécessaires pour que la struct satisfasse l'interface `Greeter` que vous avez mise en place à la tâche 1.
Vous pouvez saluer quelqu'un en italien avec `"Ciao {name}!"`.

## 3. Implémenter le portugais

Avant de terminer votre journée, vous êtes également censé terminer la fonctionnalité pour saluer les gens en portugais.

Pour cela, créez une struct `Portuguese` et implémentez les deux méthodes nécessaires pour que la struct satisfasse l'interface `Greeter` ici aussi.
Vous pouvez saluer quelqu'un en portugais avec `"Olá {name}!"`.

## Source

### Créé par

- @junedev
