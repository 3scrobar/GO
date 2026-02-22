# Les Jouets d'Elon

Bienvenue dans « Les Jouets d'Elon » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Une méthode est une fonction avec un argument de _récepteur_ spécial. Le récepteur apparaît dans sa propre liste d'arguments entre le mot-clé `func` et le nom de la méthode.

```go
func (receiver type) MethodName(parameters) (returnTypes) {
}
```

Vous ne pouvez définir une méthode avec un récepteur dont le type est défini que dans le même package que la méthode.

```go
package person 

type Person struct {
	Name string
}

func (p Person) Greetings() string {
	return fmt.Sprintf("Welcome %s!", p.Name)
}
```

La méthode sur la struct peut être appelée via la notation pointée.

```go
p := Person{Name: "Bronson"}
fmt.Println(p.Greetings())
// Output: Welcome Bronson!
```

Remarquez comment nous avons appelé la méthode `Greetings()` sur l'instance `Person` `p`.
C'est exactement comme la façon dont vous appelez les méthodes dans un langage de programmation orienté objet.

Rappelez-vous : une méthode est juste une fonction avec un argument récepteur.
Les méthodes aident à éviter les conflits de noms - puisqu'une méthode est liée à un type de récepteur particulier, vous pouvez avoir le même nom de méthode sur différents types.

```go
import "math"

type rect struct {
	width, height int
}
func (r rect) area() int {
	return r.width * r.height
}

type circle struct {
	radius int
}
func (c circle) area() float64 {
	return math.Pow(float64(c.radius), 2) * math.Pi
}
```

Il existe deux types de récepteurs : les récepteurs par valeur et les récepteurs par pointeur.

Toutes les méthodes que nous avons vues jusqu'à présent ont un récepteur par valeur, ce qui signifie qu'elles recevront une copie de la valeur transmise à la méthode, ce qui signifie que toute modification apportée au récepteur à l'intérieur de la méthode n'est pas visible par l'appelant.

Vous pouvez déclarer des méthodes avec des récepteurs par pointeur pour modifier la valeur vers laquelle pointe le récepteur.
Ceci est fait en préfixant le nom du type avec un `*`.
Par exemple avec le type `rect`, un récepteur par pointeur serait déclaré comme `*rect`.
De telles modifications sont également visibles par l'appelant de la méthode.

```go
type rect struct {
	width, height int
}
func (r *rect) squareIt() {
	r.height = r.width
}

r := rect{width: 10, height: 20}
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 20

r.squareIt()
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 10
```

## Instructions

Remarque : Cet exercice est une continuation de l'exercice `need-for-speed`.

Dans cet exercice, vous organiserez des courses entre différents types de voitures de contrôle à distance. Chaque voiture a ses propres caractéristiques de vitesse et de consommation de batterie.

Les voitures commencent avec une batterie pleine (100%). Chaque fois que vous conduisez la voiture à l'aide de la télécommande, elle parcourt la vitesse de la voiture en mètres et diminue le pourcentage de batterie restant par sa consommation de batterie.

Si la batterie d'une voiture est inférieure à son pourcentage de consommation de batterie, vous ne pouvez plus conduire la voiture.

La voiture de contrôle à distance a un affichage LED sophistiqué qui affiche deux informations :

- La distance totale parcourue, affichée comme : `"Driven <METERS> meters"`.
- La charge de batterie restante, affichée comme : `"Battery at <PERCENTAGE>%"`.

Chaque piste de course a sa propre distance. Les voitures sont testées en vérifiant si elles peuvent terminer la piste sans manquer de batterie.

## 1. Conduire la voiture

Implémentez la méthode `Drive` sur la `Car` qui met à jour le nombre de mètres parcourus en fonction de la vitesse de la voiture et réduit la batterie selon la consommation de batterie :

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car.Drive()
// car is now Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
```

Remarque : Vous ne devez pas essayer de conduire la voiture si cela entraînerait la batterie de la voiture en dessous de 0.

## 2. Afficher la distance parcourue

Implémentez une méthode `DisplayDistance` sur `Car` pour retourner la distance affichée sur l'affichage LED sous forme de `string` :

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

fmt.Println(car.DisplayDistance())
// Output: "Driven 0 meters"
```

## 3. Afficher le pourcentage de batterie

Implémentez la méthode `DisplayBattery` sur `Car` pour retourner le pourcentage de batterie affiché sur l'affichage LED sous forme de `string` :

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

fmt.Println(car.DisplayBattery())
// Output: "Battery at 100%"
```

## 4. Vérifier si une voiture de contrôle à distance peut terminer une course

Pour terminer une course, une voiture doit pouvoir parcourir la distance de la piste. Cela signifie ne pas décharger sa batterie avant d'avoir franchi la ligne d'arrivée. Implémentez la méthode `CanFinish` qui prend un `trackDistance int` en tant que paramètre et retourne `true` si la voiture peut terminer la course ; sinon, retournez `false` :

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

trackDistance := 100

car.CanFinish(trackDistance)
// => true
```

## Source

### Créé par

- @tehsphinx

### Contribué par

- @oanaOM
- @mcastorina
