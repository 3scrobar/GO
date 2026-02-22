# Besoin de Vitesse

Bienvenue dans « Besoin de Vitesse » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

En Go, une `struct` est une séquence d'éléments nommés appelés _champs_, chaque champ a un nom et un type.
Le nom d'un champ doit être unique au sein de la struct.
Les structs peuvent être comparées aux _classes_ du paradigme de programmation orientée objet.

## Définir une struct

Vous créez une nouvelle struct en utilisant les mots-clés `type` et `struct`, et en définissant explicitement le nom et le type des champs.
Par exemple, ici nous définissons la struct `Shape`, qui contient le nom d'une forme et sa taille :

```go
type Shape struct {
	name string
	size int
}
```

Les noms de champs dans les structs suivent la convention Go - les champs dont le nom commence par une lettre minuscule ne sont visibles que pour le code du même package, tandis que ceux dont le nom commence par une lettre majuscule sont visibles dans d'autres packages.

## Créer des instances d'une struct

Une fois que vous avez défini la `struct`, vous devez créer une nouvelle instance en définissant les champs en utilisant leur nom de champ dans n'importe quel ordre :

```go
s := Shape {
	name: "Square",
	size: 25,
}
```

Pour lire ou modifier des champs d'instance, utilisez la notation `.` :

```go
// Mettre à jour le nom et la taille
s.name = "Circle"
s.size = 35
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name: Circle size: 35
```

Les champs qui n'ont pas de valeur initiale assignée auront leur valeur zéro. Par exemple :

```go
s := Shape{name: "Circle"}
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name: Circle size: 0
```

Vous pouvez créer une instance d'une `struct` sans utiliser les noms de champs, tant que vous définissez les champs _dans l'ordre_ :

```go
s := Shape {
	"Oval",
	20,
}
```

Cependant, cette syntaxe est considérée comme du _code fragile_ car elle peut se casser quand un champ est ajouté, en particulier quand le nouveau champ est d'un type différent.
Dans l'exemple suivant, nous ajoutons un champ supplémentaire à `Shape` :

```go
type Shape struct {
	name        string
	description string // nouveau champ 'description' ajouté
	size        int
}

s := Shape{
	"Circle",
	20,
}
// Puisque le deuxième champ de la struct est maintenant une chaîne et non un int,
// le compilateur lèvera une erreur lors de la compilation du programme :
// Output: cannot use 20 (type untyped int) as type string in field value
// Output: too few values in Shape{...}
```

## Fonctions "Nouveau"

Parfois, il peut être agréable d'avoir des fonctions qui nous aident à créer des instances de struct.
Par convention, ces fonctions sont généralement appelées `New` ou leurs noms commencent par `New`, mais comme ce sont juste des fonctions ordinaires, vous pouvez leur donner n'importe quel nom.
Elles pourraient vous rappeler les constructeurs dans d'autres langages, mais en Go, ce sont juste des fonctions ordinaires.

Dans l'exemple suivant, l'une de ces fonctions `New` est utilisée pour créer une nouvelle instance de `Shape` et définir automatiquement une valeur par défaut pour la `taille` de la forme :

```go
func NewShape(name string) Shape {
	return Shape{
		name: name,
		size: 100, //valeur par défaut pour size est 100
	}
}

NewShape("Triangle")
// => Shape { name: "Triangle", size: 100 }

```

L'utilisation de fonctions `New` peut avoir les avantages suivants :
* validation des valeurs données
* gestion des valeurs par défaut
* comme les fonctions `New` sont souvent déclarées dans le même package que les structs qu'elles initializent, elles peuvent initialiser même les champs privés de la struct

## Instructions

Dans cet exercice, vous organiserez des courses entre différents types de voitures télécommandées.
Chaque voiture a ses propres caractéristiques de vitesse et de consommation de batterie.

Les voitures commencent avec des batteries complètes (100%). Chaque fois que vous conduisez la voiture à l'aide de la télécommande,
elle parcourt la vitesse de la voiture en mètres et diminue le pourcentage de batterie restant de sa consommation de batterie.

Si la batterie d'une voiture est inférieure à son pourcentage de consommation de batterie, vous ne pouvez plus conduire la voiture.

Chaque piste de course a sa propre distance. Les voitures sont testées en vérifiant si elles peuvent terminer la piste sans manquer de batterie.

## 1. Créer une voiture télécommandée

Définissez une struct `Car` avec les champs de type `int` suivants :

- batterie
- vidange de la batterie
-vitesse
- distance

Permettez de créer une voiture télécommandée en définissant une fonction `NewCar` qui prend la vitesse de la voiture en mètres,
et le pourcentage de consommation de batterie comme ses deux paramètres (tous deux de type `int`) et retourne une instance `Car` :

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
// => Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}
```

## 2. Créer une piste de course

Définissez un autre type de struct appelé `Track` avec le champ `distance` de type entier.
Permettez de créer une piste de course en définissant une fonction `NewTrack` qui prend la distance de la piste en mètres comme seul paramètre (de type `int`) :

```go
distance := 800
track := NewTrack(distance)
// => Track{distance: 800}
```

## 3. Conduire la voiture

Implémentez la fonction `Drive` qui met à jour le nombre de mètres parcourus en fonction de la vitesse de la voiture, et réduit la batterie selon la consommation de batterie.
Si la batterie n'est pas suffisante pour conduire une fois de plus, la voiture ne bougera pas :

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car = Drive(car)
// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
```

## 4. Vérifier si une voiture télécommandée peut terminer une course

Pour terminer une course, une voiture doit être capable de parcourir la distance de la course. Cela signifie ne pas vider sa batterie avant d'avoir franchi la ligne d'arrivée. Implémentez la fonction `CanFinish` qui prend une instance `Car` et `Track` comme paramètre et retourne `true` si la voiture peut terminer la course ; sinon, retournez `false`.

Supposez que vous êtes actuellement à la ligne de départ de la course et démarrez le moteur de la voiture pour la course. Tenez compte du fait que la batterie de la voiture pourrait ne pas être complètement chargée au démarrage de la course :
```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

distance := 100
track := NewTrack(distance)

CanFinish(car, track)
// => true
```

## Source

### Créé par

- @tehsphinx

### Contribué par

- @oanaOM
- @eklatzer
- @andrerfcsantos
