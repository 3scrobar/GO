# Échiquier

Bienvenue dans « Échiquier » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

En Go, vous pouvez itérer sur une `slice` en utilisant `for` et un index, ou vous pouvez utiliser `range`.
`range` vous permet également d'itérer sur une `map`.

Chaque itération retourne deux valeurs : l'index/la clé et une copie de l'élément à cet index/cette clé.

## Itérer sur une slice

Facile comme bonjour, boucle sur une slice, ordonnée comme prévu.

```go
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(i, x)
}
// affiche:
// 0, 10
// 1, 20
// 2, 30
```

## Itérer sur une map

Itérer sur une map pose un nouveau problème. L'ordre est maintenant aléatoire.

```go
hash := map[int]int{9: 10, 99: 20, 999: 30}
for k, v := range hash {
  fmt.Println(k, v)
}
// affiche, par exemple:
// 99 20
// 999 30
// 9 10
```

~~~~exercism/note
La sortie ci-dessus peut sembler incorrecte, car on s'attendrait à ce que la première paire clé/valeur de la déclaration de la map `9 10` soit imprimée en premier et non en dernier.
Cependant, les maps sont de nature non ordonnées - il n'y a pas de première ou dernière paire clé/valeur.
Pour cette raison, lors de l'itération sur les entrées d'une map, l'ordre dans lequel les entrées seront visitées sera aléatoire et ne suivra aucun motif spécifique.
Cela signifie que la sortie ci-dessus est possible mais peut différer de ce que vous obtiendrez si vous essayez d'exécuter cela vous-même.
Pour en savoir plus, consultez [Go Language Spec: range clause](https://go.dev/ref/spec#RangeClause).
~~~~

## Itération en omettant la clé ou la valeur

En Go, une variable inutilisée lèvera une erreur au moment de la compilation.
Parfois, vous n'avez besoin que de la valeur, comme dans le premier exemple :

```go
xi := []int{10, 20, 30}
for i, x := range xi {
  fmt.Println(x)
}
// Écec de la compilation Go: i déclaré mais non utilisé
```

Vous pouvez remplacer le `i` par `_` qui indique au compilateur que nous n'utilisons pas cette valeur :

```go
xi := []int{10, 20, 30}
for _, x := range xi {
  fmt.Println(x)
}
// affiche:
// 10
// 20
// 30
```

Si vous voulez seulement imprimer l'index, vous pouvez remplacer le `x` par `_`,
ou simplement omettre la déclaration :

```go
xi := []int{10, 20, 30}
// for i, _ := range xi {
for i := range xi {
  fmt.Println(i)
}
// affiche:
// 0
// 1
// 2
```

## Types non structurés

Vous avez déjà vu la définition de types struct.
Il est également possible de définir des types non-struct que vous pouvez utiliser comme alias pour une déclaration de type intégré, et vous pouvez définir des fonctions récepteur sur eux pour les étendre de la même manière que les types struct.

```go
type Name string
func SayHello(n Name) {
  fmt.Printf("Hello %s\n", n)
}
n := Name("Fred")
SayHello(n)
// Sortie: Hello Fred
```

Vous pouvez également définir des types non-struct composés de tableaux et de maps.

```go
type Names []string
func SayHello(n Names) {
  for _, name := range n {
    fmt.Printf("Hello %s\n", name)
  }
}
n := Names([]string{"Fred", "Bill"})
SayHello(n)
// Sortie:
// Hello Fred
// Hello Bill
```

## Instructions

En tant qu'amateur d'échecs, vous aimeriez créer votre propre version du jeu. Oui, il peut y avoir de nombreuses implémentations d'échecs disponibles en ligne, mais la vôtre sera unique !

Chaque case de l'échiquier est identifiée par une paire lettre-nombre :
- Les rangées horizontales de cases, appelées rangées, sont numérotées de 1 à 8.
- Les colonnes verticales de cases, appelées colonnes, sont étiquetées de A à H.

```
   A B C D E F G H
 8 # _ _ _ # _ _ # 8
 7 _ _ _ _ _ _ _ _ 7
 6 _ _ _ _ # _ _ # 6
 5 _ # _ _ _ _ _ # 5
 4 _ _ _ _ _ _ # # 4
 3 # _ # _ _ _ _ # 3
 2 _ _ _ _ _ _ _ # 2
 1 # _ _ _ _ _ _ # 1
   A B C D E F G H
```

## 1. Étant donné un échiquier et une colonne, comptez le nombre de cases occupées

Implémentez la fonction `CountInFile(board Chessboard, file string) int`.
Elle doit compter le nombre total de cases occupées en itérant sur une map. Retournez un entier.
Retournez un décompte de zéro (`0`) si la colonne donnée ne peut pas être trouvée dans la map.

```go
CountInFile(board, "A")
// => 3
```

## 2. Étant donné un échiquier et une rangée, comptez le nombre de cases occupées

Implémentez la fonction `CountInRank(board Chessboard, rank int) int`.
Elle doit compter le nombre total de cases occupées en itérant sur la rangée donnée. Retournez un entier.
Retournez un décompte de zéro (`0`) si la rangée donnée n'est pas valide (pas entre `1` et `8`, inclus).

```go
CountInRank(board, 2)
// => 1
```

## 3. Comptez le nombre de cases présentes dans l'échiquier donné

Implémentez la fonction `CountAll(board Chessboard) int`.
Elle doit compter combien de cases sont présentes dans l'échiquier et retourner
un entier. Puisque vous n'avez pas besoin de vérifier le contenu des cases,
envisagez d'utiliser range en omettant à la fois l'`index` et la `valeur`.

```go
CountAll(board)
// => 64
```

## 4. Comptez le nombre de cases occupées dans l'échiquier donné

Implémentez la fonction `CountOccupied(board Chessboard) int`.
Elle doit compter le nombre de cases occupées dans l'échiquier.
Retournez un entier.

```go
CountOccupied(board)
// => 15
```

## Source

### Créé par

- @brugnara
- @tehsphinx

### Contribué par

- @eklatzer
