# Jeux de Cartes

Bienvenue dans « Jeux de Cartes » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

## Tranches (Tranches)

Les tranches en Go sont similaires aux listes ou tableaux dans d'autres langages.
Elles contiennent plusieurs éléments d'un type spécifique (ou interface).

Les tranches en Go sont basées sur des tableaux.
Les tableaux ont une taille fixe.
Une tranche, en revanche, est une vue flexible et de taille dynamique des éléments d'un tableau.

Une tranche s'écrit comme `[]T` où `T` est le type des éléments de la tranche :

```go
var empty []int                 // une tranche vide
withData := []int{0,1,2,3,4,5}  // une tranche pré-remplie avec des données
```

Vous pouvez obtenir ou définir un élément à un index de base zéro donné en utilisant la notation entre crochets :

```go
withData[1] = 5
x := withData[1] // x vaut maintenant 5
```

Vous pouvez créer une nouvelle tranche à partir d'une tranche existante en obtenant une plage d'éléments.
Encore une fois en utilisant la notation entre crochets, mais en spécifiant à la fois un index de départ (inclus) et un index de fin (exclus).
Si vous ne spécifiez pas un index de départ, il vaut par défaut 0.
Si vous ne spécifiez pas un index de fin, il vaut par défaut la longueur de la tranche.

```go
newSlice := withData[2:4]
// => []int{2,3}
newSlice := withData[:2]
// => []int{0,1}
newSlice := withData[2:]
// => []int{2,3,4,5}
newSlice := withData[:]
// => []int{0,1,2,3,4,5}
```

Vous pouvez ajouter des éléments à une tranche en utilisant la fonction `append`.
Ci-dessous, nous ajoutons `4` et `2` à la tranche `a`.

```go
a := []int{1, 3}
a = append(a, 4, 2)
// => []int{1,3,4,2}
```

`append` retourne toujours une nouvelle tranche, et quand nous voulons simplement ajouter des éléments à une tranche existante, il est courant de la réassigner à la variable de tranche que nous avons passée comme premier argument, comme nous l'avons fait ci-dessus.

`append` peut également être utilisé pour fusionner deux tranches :

```go
nextSlice := []int{100,101,102}
newSlice  := append(withData, nextSlice...)
// => []int{0,1,2,3,4,5,100,101,102}
```

## Fonctions Variadiques

Habituellement, les fonctions en Go n'acceptent qu'un nombre fixe d'arguments.
Cependant, il est également possible d'écrire des fonctions variadiques en Go.

Une fonction variadique est une fonction qui accepte un nombre variable d'arguments.

Si le type du dernier paramètre dans une définition de fonction est préfixé par une ellipse `...`, alors la fonction peut accepter un nombre quelconque d'arguments pour ce paramètre.

```go
func find(a int, b ...int) {
    // ...
}
```

Dans la fonction ci-dessus, le paramètre `b` est variadique et nous pouvons passer 0 ou plusieurs arguments à `b`.

```go
find(5, 6)
find(5, 6, 7)
find(5)
```

~~~~exercism/caution
Le paramètre variadique doit être le dernier paramètre de la fonction.
~~~~

La façon dont les fonctions variadiques fonctionnent est en convertissant le nombre variable d'arguments en une tranche du type du paramètre variadique.

Voici un exemple d'implémentation d'une fonction variadique.

```go
func find(num int, nums ...int) {
    fmt.Printf("type of nums is %T\n", nums)

    for i, v := range nums {
        if v == num {
            fmt.Println(num, "found at index", i, "in", nums)
            return
        }
    }

    fmt.Println(num, "not found in ", nums)
}

func main() {
    find(89, 90, 91, 95)
    // =>
    // type of nums is []int
    // 89 not found in  [90 91 95]

    find(45, 56, 67, 45, 90, 109)
    // =>
    // type of nums is []int
    // 45 found at index 2 in [56 67 45 90 109]

    find(87)
    // =>
    // type of nums is []int
    // 87 not found in  []
}
```

À la ligne `find(89, 90, 91, 95)` du programme ci-dessus, le nombre variable d'arguments à la fonction find sont `90`, `91` et `95`.
La fonction `find` s'attend à un paramètre int variadique après `num`.
Par conséquent, ces trois arguments seront convertis par le compilateur en une tranche de type `int` `[]int{90, 91, 95}` puis seront passés à la fonction find en tant que `nums`.

Parfois, vous avez déjà une tranche et vous voulez la passer à une fonction variadique.
Ceci peut être réalisé en passant la tranche suivie de `...`.
Cela indiquera au compilateur d'utiliser la tranche telle quelle à l'intérieur de la fonction variadique.
L'étape décrite ci-dessus où une tranche est créée sera simplement omise dans ce cas.

```go
list := []int{1, 2, 3}
find(1, list...) // "find" définie comme indiqué ci-dessus
```

[append-yourbasic]: https://yourbasic.org/golang/append-explained/

## Instructions

En tant que magicienne en herbe, Elyse doit pratiquer les bases. Elle a un jeu de cartes qu'elle veut manipuler.

Pour simplifier les choses, elle n'utilise que les cartes 1 à 10.

## 1. Créer une tranche avec certaines cartes

Lorsqu'elle s'entraîne avec ses cartes, Elyse aime commencer avec ses trois cartes préférées du jeu : 2, 6 et 9.
Écrivez une fonction `FavoriteCards` qui retourne une tranche avec ces cartes dans cet ordre.

```go
cards := FavoriteCards()
fmt.Println(cards)
// Sortie: [2 6 9]
```

## 2. Récupérer une carte d'une pile

Retournez la carte à la position `index` de la pile donnée.

```go
card := GetItem([]int{1, 2, 4, 1}, 2) // card == 4
```

Si l'index est hors limites (c.-à-d. s'il est négatif ou après la fin de la pile), nous voulons retourner `-1` :

```go
card := GetItem([]int{1, 2, 4, 1}, 10) // card == -1
```

~~~~exercism/note
Par convention en Go, une erreur est retournée au lieu de retourner une valeur "hors-bande".
Ici, la valeur "hors-bande" est `-1` quand un entier positif est attendu.
Lors du retour d'une erreur, il est considéré idiomatique de retourner la [`valeur zéro`](https://www.geeksforgeeks.org/zero-value-in-golang/) avec l'erreur.
Le retour d'une erreur avec la valeur appropriée sera couvert dans un exercice futur.
~~~~

## 3. Échanger une carte dans la pile

Échangez la carte à la position `index` avec la nouvelle carte fournie et retournez la pile ajustée.
Notez que cela modifiera la tranche d'entrée, ce qui est le comportement attendu.

```go
index := 2
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards)
// Sortie: [1 2 6 1]
```

Si l'index est hors limites (c.-à-d. s'il est négatif ou après la fin de la pile), nous voulons ajouter la nouvelle carte à la fin de la pile :

```go
index := -1
newCard := 6
cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
fmt.Println(cards)
// Sortie: [1 2 4 1 6]
```

## 4. Ajouter des cartes au début de la pile

Ajoutez la (les) carte(s) spécifiée(s) dans le paramètre `value` au début de la pile.

```go
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice, 5, 1)
fmt.Println(cards)
// Sortie: [5 1 3 2 6 4 8]
```

Si aucun argument n'est donné pour le paramètre `value`, alors le résultat correspond à la tranche originale.

```go
slice := []int{3, 2, 6, 4, 8}
cards := PrependItems(slice)
fmt.Println(cards)
// Sortie: [3 2 6 4 8]
```

## 5. Supprimer une carte de la pile

Supprimez la carte à la position `index` de la pile et retournez la pile.
Notez que cela peut modifier la tranche d'entrée, ce qui est acceptable.

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
fmt.Println(cards)
// Sortie: [3 2 4 8]
```

Si l'index est hors limites (c.-à-d. s'il est négatif ou après la fin de la pile), nous voulons laisser la pile inchangée :

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
fmt.Println(cards)
// Sortie: [3 2 6 4 8]
```

## Source

### Créé par

- @tehsphinx

### Contribué par

- @norbs57
