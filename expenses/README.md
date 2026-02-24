# Dépenses

Bienvenue dans Dépenses sur la piste Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez de le résoudre sans les utiliser d'abord :)

## Introduction

En Go, les fonctions sont des valeurs de première classe. Cela signifie que vous pouvez faire avec les fonctions les mêmes choses que vous pouvez faire avec toutes les autres valeurs - assigner des fonctions à des variables, les passer en tant qu'arguments à d'autres fonctions ou même retourner des fonctions à partir d'autres fonctions.

Ci-dessous, nous créons deux fonctions, `engGreeting` et `espGreeting` et nous les assignons à la variable `greeting` :

```go
import "fmt"

func engGreeting(name string) string {
  return fmt.Sprintf("Hello %s, nice to meet you!", name)
}

func espGreeting(name string) string {
  return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := engGreeting			// greeting est une variable de type func(string) string
fmt.Println(greeting("Alice"))	// Hello Alice, nice to meet you!

greeting = espGreeting
fmt.Println(greeting("Alice")) 	// ¡Hola Alice, mucho gusto!
```

Les valeurs de fonction offrent une opportunité de paramétrer les fonctions non seulement avec des données mais aussi avec un comportement.
Dans l'exemple suivant, nous passons un comportement à la fonction `dialog` via le paramètre `greetingFunc` :

```go
func dialog(name string, greetingFunc func(string) string) {
  fmt.Println(greetingFunc(name))
  fmt.Println("I'm a dialog bot.")
}

func espGreeting(name string) string {
  return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := espGreeting
dialog("Alice", greeting)
// =>
// ¡Hola Alice, mucho gusto!
// I'm a dialog bot.
```

La valeur d'une variable non initialisée de type fonction est `nil`.
Par conséquent, appeler une valeur de fonction `nil` provoque une panique.

```go
var dutchGreeting func(string) string
dutchGreeting("Alice") // panic: call of nil function
```

Les valeurs de fonction peuvent être comparées avec `nil`. Cela peut être utile pour éviter des paniques de programme inutiles.

```go
var dutchGreeting func(string) string
if dutchGreeting != nil {
  dutchGreeting("Alice") // safe to call dutchGreeting
}
```

## Types de fonction

L'utilisation de valeurs de fonction est possible grâce aux types de fonction en Go. Un type de fonction désigne l'ensemble de toutes les fonctions ayant la même séquence de types de paramètres et la même séquence de types de résultats. Les types définis par l'utilisateur peuvent être déclarés au-dessus des types de fonction. Par exemple, la fonction `dialog` des exemples précédents peut être mise à jour comme suit :

```go
type greetingFunc func(string) string

func dialog(name string, f greetingFunc) {
  fmt.Println(f(name))
  fmt.Println("I'm a dialog bot.")
}
```

## Fonctions anonymes

Un autre outil puissant disponible grâce au support des fonctions de première classe sont les fonctions anonymes. Les fonctions anonymes sont définies au point d'utilisation, sans nom suivant le mot-clé `func`. Ces fonctions ont accès aux variables de la fonction englobante.

Par exemple :

```go
func fib() func() int {
  var n1, n2 int

  return func() int {
    if n1 == 0 && n2 == 0 {
      n1 = 1
    } else {
      n1, n2 = n2, n1 + n2
    }
    return n2
  }
}

next := fib()
for i := 0; i < N; i++ {
  fmt.Printf("F%d\t= %4d\n", i, next())
}
```

Un appel à `fib` déclare les variables `n1` et `n2` et retourne une fonction anonyme qui, à son tour, change les valeurs de ces variables à chaque appel de la fonction. Les appels Nth de la fonction anonyme retournent le Nième nombre de la séquence de Fibonacci à partir de 0. La fonction anonyme interne a accès aux variables locales (`n1` et `n2`) de la fonction englobante `fib`. C'est un excellent moyen de faire en sorte que les valeurs de fonction conservent l'état entre les appels. On dit que la fonction anonyme est une fermeture des variables `n1` et `n2`. [Les fermetures][closure] sont largement utilisées en programmation et vous pourriez les voir supportées par d'autres langages.

[closure]: https://en.wikipedia.org/wiki/Closure_(computer_programming)

## Instructions

Bob est un conseiller financier et aide les gens à gérer leurs dépenses. Les clients de Bob envoient des enregistrements de dépenses pour qu'il les analyse. Bob a des enregistrements des périodes précédentes pour pouvoir voir les changements dans les dépenses. Bob n'aime pas les calendriers et utilise l'*époque de Bob* à la place des dates. L'époque de Bob est le nombre de jours écoulés depuis que le client de Bob a commencé son activité.

Dans cet exercice, vous allez construire un programme pour aider Bob à gérer et analyser les dépenses de ses clients.

Bob travaille avec des `Record` et des `DaysPeriod`.

Un `Record` représente un enregistrement de dépenses contenant le jour auquel la dépense a été effectuée, l'argent dépensé et la catégorie de la dépense.

```go
// Record représente un enregistrement de dépenses.
type Record struct {
  Day         int
  Amount      float64
  Category    string
}
```

Un `DaysPeriod` représente une plage de jours et inclut tous les jours du jour `From` jusqu'au jour `To`.
Les deux extrémités sont incluses dans la plage.

```go
// DaysPeriod représente une période de jours.
type DaysPeriod struct {
  From int
  To   int
}

p := DaysPeriod{From: 1, To: 31}
// p représente tous les jours du jour 1 au jour 31:
//  - les jours 1, 20, 16 et 31 sont des exemples de jours inclus
//    dans la plage de temps spécifiée par p
//  - les jours 50 et 40 sont des exemples de jours non inclus
//    dans la plage de temps spécifiée par p  
```

## 1. Implémenter un filtre d'enregistrements général

Bob traite beaucoup d'enregistrements chaque jour, mais tous ne sont pas intéressants selon l'analyse que Bob effectue.
Aidons Bob à effectuer un filtrage basique des enregistrements.

Implémentez la fonction `Filter` générique pour filtrer les enregistrements selon un critère donné par une fonction.
Cette fonction de filtre accepte une collection d'enregistrements et une fonction de prédicat et retourne uniquement les enregistrements de la collection qui satisfont le prédicat.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
}

// Day1Records retourne vrai uniquement pour les enregistrements du jour 1
func Day1Records(r Record) bool {
  return r.Day == 1
}

Filter(records, Day1Records)
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"}
// ]
```

## 2. Filtrer les enregistrements dans une période de temps

Bob a fréquemment besoin de filtrer les enregistrements qui se situent dans une période de temps donnée.

Implémentez la fonction `ByDaysPeriod` qui aidera Bob à créer de tels filtres.
Cette fonction accepte un `DaysPeriod` et retourne une fonction qui prend un enregistrement et indique si l'enregistrement se situe dans la période de temps spécifiée par le `DaysPeriod` donné en argument.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

period := DaysPeriod{From: 1, To: 15}

Filter(records, ByDaysPeriod(period))
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 11, Amount: 300, Category: "utility-bills"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]
```

## 3. Filtrer les enregistrements par catégorie

Au-delà du filtrage des enregistrements par période de temps, Bob a également besoin de filtrer les enregistrements par catégorie.

Implémentez la fonction `ByCategory` qui aidera Bob à créer de tels filtres.
Cette fonction accepte une catégorie et retourne une fonction qui prend un enregistrement et indique si la catégorie de cet enregistrement est la même que la catégorie donnée en argument.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

Filter(records, ByCategory("groceries"))
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"},
//   {Day: 12, Amount: 28, Category: "groceries"},
// ]
```

## 4. Calculer le montant total des dépenses dans une période

Bob veut aussi connaître le montant total des dépenses dans une période de temps.

Implémentez la fonction `TotalByPeriod` pour retourner une somme des dépenses dans la plage de jours.

```go
records := []Record{
  {Day: 15, Amount: 16, Category: "entertainment"},
  {Day: 32, Amount: 20, Category: "groceries"},
  {Day: 40, Amount: 30, Category: "entertainment"}
}

p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

TotalByPeriod(records, p1)
// => 16

TotalByPeriod(records, p2)
// => 50
```

## 5. Calculer les dépenses totales pour les enregistrements d'une catégorie dans une période

Pour les rapports les plus complexes que Bob fait à ses clients, Bob a besoin de filtrer les enregistrements par catégorie et période de temps en même temps.
Cela signifie que Bob veut connaître les dépenses totales pour les enregistrements dans une catégorie dans une période de temps donnée.

Implémentez la fonction `CategoryExpenses` qui retourne le montant total des dépenses dans une catégorie dans une période de jours donnée. La fonction doit également différencier le cas où la catégorie donnée n'est pas présente dans les enregistrements de dépenses du cas où il n'y a pas de dépenses de la catégorie dans la période fournie.
Lorsque la catégorie n'est la catégorie d'aucun des enregistrements (indépendamment de la période), la fonction doit retourner une erreur.

```go
p1 := DaysPeriod{From: 1, To: 30}
p2 := DaysPeriod{From: 31, To: 60}

records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
  {Day: 26, Amount: 300, Category: "university"},
  {Day: 28, Amount: 1300, Category: "rent"},
}

CategoryExpenses(records, p1, "entertainment")
// => 0, error(unknown category entertainment)

CategoryExpenses(records, p1, "rent")
// => 1300, nil

CategoryExpenses(records, p2, "rent")
// => 0, nil
```

## Source

### Créé par

- @antklim
- @andrerfcsantos