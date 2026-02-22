# Magasin en Gros

Bienvenue dans « Magasin en Gros » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

En Go, `map` est un type de données intégré qui mappe les clés aux valeurs. Dans d'autres langages de programmation, vous pourriez être familier avec le concept de `map` comme dictionnaire, table de hachage, magasin clé/valeur ou tableau associatif.

Syntaxiquement, `map` ressemble à ceci :

```go
map[KeyType]ElementType
```

Il est aussi important de savoir que chaque clé est unique, ce qui signifie que l'assignation de la même clé deux fois écrasera la valeur de la clé correspondante.

Pour créer une map, vous pouvez faire :

```go
  // Avec un littéral map
  foo := map[string]int{}
```

ou

```go
  // ou avec la fonction make
  foo := make(map[string]int)
```

Voici quelques opérations que vous pouvez faire avec une map

```go
  // Ajouter une valeur dans une map avec l'opérateur `=` :
  foo["bar"] = 42
  // Ici, nous mettons à jour l'élément de `bar`
  foo["bar"] = 73
  // Pour récupérer une valeur de map, vous pouvez utiliser
  baz := foo["bar"]
  // Pour supprimer un élément d'une map, vous pouvez utiliser
  delete(foo, "bar")
```

Si vous essayez de récupérer la valeur d'une clé qui n'existe pas dans la map, elle renverra la valeur zéro du type de valeur.
Cela peut vous confondre, surtout si la valeur par défaut de votre `ElementType` (par exemple, 0 pour un int), est une valeur valide.
Pour vérifier si une clé existe dans votre map, vous pouvez utiliser

```go
  value, exists := foo["baz"]
  // Si la clé "baz" n'existe pas,
  // value: 0; exists: false
```

## Instructions

Un ami à vous possède un ancien magasin de gros appelé **Magasin en Gros**.
Le nom vient de la quantité d'articles que le magasin vend : tout en [unité de gros][gross-unit].
Votre ami vous a demandé d'implémenter un système de point de vente (POS) pour son magasin.
**D'abord, vous voulez construire un prototype pour cela.**
**Dans votre prototype, votre système enregistrera uniquement la quantité.**
Votre ami vous a donné une liste de mesures pour vous aider :

| Unité              | Score |
| :------------------- | ----  : |
| quart_de_douzaine | 3 |
| une demi-douzaine | 6 |
| douzaine | 12 |
| petit_gross | 120 |
| brute | 144 |
| super_gross | 1728 |

## 1. Stocker l'unité de mesure dans votre programme

Pour utiliser la mesure, vous devez stocker la mesure dans votre programme.

```go
units := Units()
fmt.Println(units)
// Résultat : map[...] avec des entrées comme ("dozen": 12)
```

## 2. Créer une nouvelle facture client

Vous devez implémenter une fonction qui crée une nouvelle facture (vide) pour le client.

```go
bill := NewBill()
fmt.Println(bill)
// Résultat : map[]
```

## 3. Ajouter un article à la facture du client

Pour implémenter cela, vous devez :

- Retourner `false` si l'`unit` donnée n'est pas dans la map `units`.
- Sinon, ajouter l'article à la `bill` du client, indexé par le nom de l'article, puis retourner `true`.
- Si l'article est déjà présent dans la facture, augmentez sa quantité du montant qui appartient à l'`unit` fournie.

```go
bill := NewBill()
units := Units()
ok := AddItem(bill, units, "carrot", "dozen")
fmt.Println(ok)
// Résultat : true (puisque dozen est une unité valide)
```

> Notez que la valeur renvoyée est de type `bool`.

## 4. Supprimer un article de la facture du client

Pour implémenter cela, vous devez :

- Retourner `false` si l'article donné n'est **pas** dans la facture
- Retourner `false` si l'`unit` donnée n'est pas dans la map `units`.
- Retourner `false` si la nouvelle quantité serait inférieure à 0.
- Si la nouvelle quantité est 0, supprimer complètement l'article de la `bill` puis retourner `true`.
- Sinon, réduire la quantité de l'article et retourner `true`.

```go
bill := NewBill()
units := Units()
ok := RemoveItem(bill, units, "carrot", "dozen")
fmt.Println(ok)
// Résultat : false (car il n'y a pas de carottes dans la facture)
```

> Notez que la valeur renvoyée est de type `bool`.

## 5. Retourner la quantité d'un article spécifique qui se trouve dans la facture du client

Pour implémenter cela, vous devez :

- Retourner `0` et `false` si l'`item` n'est pas dans la facture.
- Sinon, retourner la quantité de l'article dans la `bill` et `true`.

```go
bill := map[string]int{"carrot": 12, "grapes": 3}
qty, ok := GetItem(bill, "carrot")
fmt.Println(qty)
// Résultat : 12
fmt.Println(ok)
// Résultat : true
```

> Notez que les valeurs renvoyées sont de types `int` et `bool`.

[gross-unit]: https://en.wikipedia.org/wiki/Gross_(unit)

## Source

### Créé par

- @chocopowwwa

### Contributions de

- @MiroslavGatsanoga
