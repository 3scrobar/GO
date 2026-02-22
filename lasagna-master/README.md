# Maître des Lasagnes

Bienvenue dans « Maître des Lasagnes » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Une fonction vous permet de regrouper du code dans une unité réutilisable.
Elle se compose du mot-clé `func`, du nom de la fonction, et d'une liste de zéro ou plusieurs paramètres typés séparés par des virgules entre parenthèses.

## Paramètres de fonction

Tous les paramètres doivent être typés explicitement ; il n'y a pas d'inférence de type pour les paramètres.
Il n'y a pas de valeurs par défaut pour les paramètres, donc tous les paramètres sont requis.

```go
import "fmt"

// Aucun paramètre
func PrintHello() {
    fmt.Println("Hello")
}

// Deux paramètres
func PrintGreetingName(greeting string, name string) {
  fmt.Println(greeting + " " + name)
}
```

Les paramètres de même type peuvent être déclarés ensemble, suivis d'une seule déclaration de type.

```go
import "fmt"

func PrintGreetingName(greeting, name string) {
  fmt.Println(greeting + " " + name)
}
```

## Paramètres vs. Arguments

Couvrons rapidement deux termes souvent confondus : `parameters` et `arguments`.
Les paramètres de fonction sont les noms définis dans la signature de la fonction, comme `greeting` et `name` dans la fonction `PrintGreetingName` ci-dessus.
Les arguments de fonction sont les valeurs concrètes passées aux paramètres quand on appelle la fonction.
Par exemple, dans l'exemple ci-dessous, `"Hello"` et `"Katrina"` sont les arguments passés aux paramètres `greeting` et `name` :

```go
PrintGreetingName("Hello", "Katrina")
```

## Valeurs de retour

Les paramètres de fonction sont suivis de zéro ou plusieurs valeurs de retour, qui doivent elles aussi être typées explicitement.
Une seule valeur de retour est indiquée sans parenthèses, plusieurs valeurs de retour sont entourées de parenthèses.
Les valeurs sont renvoyées au code appelant via le mot-clé `return`.
Il peut y avoir plusieurs instructions `return` dans une fonction.
L'exécution de la fonction se termine dès qu'elle rencontre l'une de ces instructions `return`.
Si plusieurs valeurs doivent être renvoyées, elles sont séparées par des virgules.

```go
func Hello(name string) string {
  return "Hello " + name
}

func HelloAndGoodbye(name string) (string, string) {
  return "Hello " + name, "Goodbye " + name
}
```

## Appel de fonctions

Pour appeler une fonction, on indique son nom puis on passe un argument pour chacun de ses paramètres entre parenthèses.

```go
import "fmt"

// Aucun paramètre, aucune valeur de retour
func PrintHello() {
    fmt.Println("Hello")
}
// S'appelle comme ceci :
PrintHello()

// Un paramètre, une valeur de retour
func Hello(name string) string {
  return "Hello " + name
}
// S'appelle comme ceci :
greeting := Hello("Dave")

// Plusieurs paramètres, plusieurs valeurs de retour
func SumAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
// S'appelle comme ceci :
aplusb, atimesb := SumAndMultiply(a, b)
```

## Valeurs de retour nommées et retour nu

Tout comme les paramètres, les valeurs de retour peuvent aussi être nommées.
Si des valeurs de retour nommées sont utilisées, une instruction `return` sans arguments renverra ces valeurs.
C'est ce qu'on appelle un retour « nu » (`naked return`).

```go
func SumAndMultiplyThenMinus(a, b, c int) (sum, mult int) {
    sum, mult = a+b, a*b
    sum -= c
    mult -= c
    return
}
```

## Passage par valeur vs. passage par référence

Il est aussi important de clarifier les notions de passage par valeur et de passage par référence.

Commençons par le passage par valeur.
Dans l'exemple ci-dessous, lorsque nous passons la variable `val` à la fonction `MultiplyByTwo`, nous passons une copie de `val`.
À cause de cela, `newVal` a la valeur mise à jour `4`, mais la variable originale `val` reste à `2`.
En coulisses, Go fait essentiellement une copie de la valeur d'origine pour que seule cette copie, c'est-à-dire `v`, soit modifiée par la fonction.

```go
val := 2
func MultiplyByTwo(v int) int {
    v = v * 2
    return v
}
newVal := MultiplyByTwo(val)
// newVal vaut 4, val vaut toujours 2 car seule une copie de sa valeur a été passée à la fonction
```

À strictement parler, en Go tous les arguments sont passés par valeur, c'est-à-dire qu'une copie de la valeur ou des données fournies à la fonction est créée.
Mais si vous ne voulez pas copier les données passées à une fonction et souhaitez modifier ces données dans la fonction,
alors vous devez utiliser des `pointers` (pointeurs) comme arguments, autrement dit passer par référence.

## Pointeurs

Nous utilisons un `pointer` (pointeur) pour obtenir un passage par référence.
En passant des arguments de type pointeur à une fonction,
nous pouvons modifier les données sous-jacentes passées à la fonction au lieu d'opérer uniquement sur une copie.

Pour l'instant, il suffit de savoir qu'un type pointeur se reconnaît au `*` placé devant le type dans la signature de la fonction.

```go
func HandlePointers(x, y *int) {
    // Logique pour manipuler les valeurs entières référencées par les pointeurs x et y
}
```

Si le concept de `pointer` est confus, pas de souci.
Nous avons une section dédiée plus loin dans le syllabus pour vous aider à maîtriser les pointeurs.

##Exceptions

Notez que les `slices` et les `maps` font exception à la règle mentionnée ci-dessus.
Quand nous passons une `slice` ou une `map` en argument d'une fonction, elles sont traitées comme des types pointeurs même sans `*` explicite dans le type.
Cela signifie que si nous passons une slice ou une map à une fonction et modifions ses données sous-jacentes,
les changements seront visibles sur la slice ou la map d'origine.

## Instructions

Dans cet exercice, vous allez écrire du code supplémentaire lié à la préparation et à la cuisson de vos brillantes lasagnes issues de votre livre de cuisine préféré.

Vous avez quatre tâches.
La première concerne la cuisson elle-même, les trois autres la préparation parfaite.

## 1. Estimer le temps de préparation

Pour la prochaine lasagne que vous allez préparer, vous voulez vous assurer d'avoir suffisamment de temps pour profiter de la cuisine.
Vous avez déjà planifié les couches de votre lasagne.
Vous voulez maintenant estimer la durée de préparation en fonction de ces couches.

Implémentez une fonction `PreparationTime` qui accepte une slice de couches en `[]string` et le temps moyen de préparation par couche en minutes sous forme d'`int`.
La fonction doit renvoyer l'estimation du temps total de préparation, en `int`, en se basant sur le nombre de couches.
Go n'a pas de valeurs par défaut pour les fonctions.
Si le temps moyen de préparation vaut `0` (valeur initiale par défaut d'un `int`), la valeur par défaut `2` doit être utilisée.

```go
layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
PreparationTime(layers, 3)
// => 18
PreparationTime(layers, 0)
// => 12
```

## 2. Calculer les quantités de nouilles et de sauce nécessaires

En plus du temps, vous voulez vous assurer d'avoir assez de sauce et de nouilles pour cuisiner la lasagne de vos rêves.
Pour chaque couche de nouilles dans votre lasagne, il vous faut 50 grammes de nouilles.
Pour chaque couche de sauce, il vous faut 0,2 litre de sauce.

Définissez la fonction `Quantities` qui prend une slice de couches en paramètre, de type `[]string`.
La fonction doit déterminer la quantité de nouilles et de sauce nécessaire pour préparer votre plat.
Le résultat doit être renvoyé sous deux valeurs : `noodles` en `int` et `sauce` en `float64`.

```go
Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
// => 100, 0.4
```

## 3. Ajouter l'ingrédient secret

Il y a quelque temps, vous avez rendu visite à un ami et mangé des lasagnes chez lui.
Elles étaient incroyables et contenaient quelque chose de spécial.
Votre ami vous a envoyé la liste des ingrédients et vous a dit que le dernier élément de la liste est « l'ingrédient secret » qui rendait ce plat si spécial.
Vous voulez maintenant ajouter cet ingrédient secret à votre recette.

Écrivez une fonction `AddSecretIngredient` qui accepte deux slices d'ingrédients de type `[]string` en paramètres.
Le premier paramètre est la liste envoyée par votre ami, le second est votre propre liste d'ingrédients.
Le dernier élément de votre liste vaut toujours `"?"`. La fonction doit le remplacer par le dernier élément de la liste de votre ami.
**Remarque :** `AddSecretIngredient` ne renvoie rien : vous devez modifier directement votre liste d'ingrédients.
La liste d'ingrédients de votre ami ne doit **pas** être modifiée.
Aussi, comme une `slice` est passée à une fonction comme un pointeur, les modifications apportées aux deux arguments `[]string` passés à `AddSecretIngredient` seront appliquées directement.

```go
friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}

AddSecretIngredient(friendsList, myList)
// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}
```

## 4. Mettre la recette à l'échelle

Les quantités listées dans votre livre de cuisine ne suffisent que pour deux portions.
Puisque vous souhaitez cuisiner pour plus de personnes la prochaine fois, vous voulez calculer les quantités pour un nombre de portions différent.

Implémentez une fonction `ScaleRecipe` qui prend deux paramètres.

- Une slice de quantités `float64` nécessaires pour 2 portions.
- Le nombre de portions que vous voulez préparer.

La fonction doit renvoyer une slice de `float64` contenant les quantités nécessaires pour le nombre de portions souhaité.
Vous voulez toutefois conserver la recette d'origine.
Cela signifie que l'argument `quantities` ne doit pas être modifié dans cette fonction.

```go
quantities := []float64{ 1.2, 3.6, 10.5 }
scaledQuantities := ScaleRecipe(quantities, 4)
// => []float64{ 2.4, 7.2, 21 }
```

## Source

### Créé par

- @bobtfish

### Contributions

- @sougat818
