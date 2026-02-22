# Référence Go - Version Débutant (Exercism)

Ce document est un mémo Go orienté débutant.
Il regroupe les notions des README Exercism et explique chaque commande/outil de base avec des exemples simples.

## Comment utiliser ce document

- Si tu bloques sur une ligne de code, va d'abord à l'`Index rapide`.
- Lis ensuite la section complète avec l'exemple.
- Reviens coder, puis valide avec les tests.

---

## Index rapide (notion -> section)

- `package`, `import` -> section 1
- `var`, `const`, `:=` -> section 2
- `if`, `else`, comparaisons -> section 3
- `switch` -> section 4
- `for`, `range` -> section 5
- Fonctions (`func`, paramètres, retours) -> section 6
- `fmt.Print`, `fmt.Sprintf`, `fmt.Println` -> section 7
- Conversions (`int`, `float64`, `strconv`) -> section 8
- Strings, runes, UTF-8 -> section 9
- Slices (`[]T`, `append`, suppression) -> section 10
- Maps (`map[K]V`, `delete`, `ok`) -> section 11
- Structs -> section 12
- Méthodes -> section 13
- Pointeurs (`&`, `*`) -> section 14
- Interfaces -> section 15
- `any`, assertions de type, type switch -> section 16
- Erreurs (`error`, `errors.New`, custom error) -> section 17
- Dates (`time`) -> section 18
- Documentation Go -> section 19
- Recettes rapides copiables -> section 20
- Pièges ultra fréquents -> section 21

---

## 1. `package` et `import`

### `package`

Chaque fichier Go appartient à un package.

```go
package main
```

- `main` est le package exécutable.
- Dans un exercice, tu auras souvent `package lasagna`, `package speed`, etc.

### `import`

Permet d'utiliser des packages externes ou standards.

```go
import "fmt"
```

Si tu importes plusieurs packages:

```go
import (
	"fmt"
	"strconv"
)
```

Important:
- Un import non utilisé provoque une erreur de compilation.

---

## 2. Variables et constantes

### `var`

```go
var age int = 25
var score int
```

- `score` prend la valeur zéro (`0`) automatiquement.

### `:=` (déclaration courte)

```go
name := "Alice"
```

- Très utilisé dans les fonctions.
- Ne fonctionne pas au niveau package (en dehors des fonctions).

### `const`

```go
const minutesPerLayer = 2
```

- Valeur immuable.

### Valeurs zéro (à retenir)

- `int` -> `0`
- `float64` -> `0`
- `string` -> `""`
- `bool` -> `false`
- `slice`, `map`, `pointer`, `interface` -> `nil`

---

## 3. Conditions: `if`, `else if`, `else`

```go
if battery <= 0 {
	return "empty"
} else if battery < 20 {
	return "low"
} else {
	return "ok"
}
```

Comparateurs:
- `==` égal
- `!=` différent
- `<`, `<=`, `>`, `>=`

Astuce de style:
- Retourne tôt les cas d'erreur, le code est plus lisible.

---

## 4. `switch`

Quand tu as beaucoup de cas, `switch` est souvent plus clair que plusieurs `if`.

```go
switch card {
case "ace":
	return 11
case "king", "queen", "jack":
	return 10
default:
	return 0
}
```

`default` = cas fallback si rien ne matche.

---

## 5. Boucles `for` et `range`

### Boucle classique

```go
for i := 0; i < len(values); i++ {
	fmt.Println(values[i])
}
```

### `range` sur une slice

```go
for i, v := range values {
	fmt.Println(i, v)
}
```

- `i` = index
- `v` = valeur

Si tu ne veux pas l'index:

```go
for _, v := range values {
	fmt.Println(v)
}
```

### `range` sur une map

```go
for k, v := range m {
	fmt.Println(k, v)
}
```

Important:
- L'ordre d'itération d'une map n'est pas garanti.

---

## 6. Fonctions (`func`)

### Fonction simple

```go
func Add(a int, b int) int {
	return a + b
}
```

Tu peux grouper les types identiques:

```go
func Add(a, b int) int {
	return a + b
}
```

### Retour multiple

```go
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```

Convention Go:
- Si plusieurs retours, `error` est le dernier.

### Named return (optionnel)

```go
func Stats(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b
	return
}
```

---

## 7. `fmt`: afficher et formater

### `fmt.Print`

Affiche tel quel.

```go
fmt.Print("Hello")
```

### `fmt.Println`

Affiche + ajoute un retour à la ligne.

```go
fmt.Println("Hello")
```

### `fmt.Sprintf`

Construit une string formatée (n'affiche pas).

```go
msg := fmt.Sprintf("Battery at %d%%", 87)
```

Spécificateurs utiles:
- `%d` entier
- `%f` float
- `%.1f` float avec 1 décimale
- `%s` string
- `%%` caractère `%`

Piège courant:
- `fmt.Sprint("%d", x)` n'interprète pas `%d`.
- Il faut `fmt.Sprintf`.

---

## 8. Conversions de types

Go ne fait pas de conversion implicite entre types numériques différents.

```go
i := 3
f := float64(i)
```

### `strconv`

- string -> int:

```go
n, err := strconv.Atoi("42")
```

- int -> string:

```go
s := strconv.Itoa(42)
```

Toujours vérifier `err` avec `Atoi`, `ParseFloat`, etc.

---

## 9. String, runes, UTF-8

### String = octets UTF-8

```go
s := "été"
fmt.Println(len(s)) // octets, pas forcément caractères
```

### Rune = caractère Unicode (point de code)

```go
for _, r := range s {
	fmt.Printf("%c\n", r)
}
```

### Compter les caractères visibles (Unicode)

```go
count := utf8.RuneCountInString(s)
```

### Remplacer un caractère

```go
s = strings.ReplaceAll(s, string('é'), string('e'))
```

Piège fréquent:
- Utiliser `len(s)` pour une limite UI avec emojis.

---

## 10. Slices (`[]T`)

### Créer

```go
var a []int          // nil slice
b := []int{1, 2, 3} // slice initialisée
```

### Lire et modifier

```go
x := b[1]   // lit
b[1] = 99   // modifie
```

### Découper (slicing)

```go
sub := b[1:3] // index 1 inclus, 3 exclu
```

### Ajouter avec `append`

```go
b = append(b, 4)
b = append(b, 5, 6)
```

Très important:
- `append` retourne une nouvelle slice.
- Toujours réassigner: `b = append(...)`

### Ajouter au début

```go
b = append([]int{0}, b...)
```

### Supprimer un élément d'index `i`

```go
b = append(b[:i], b[i+1:]...)
```

### Vérification d'index

```go
if i < 0 || i >= len(b) {
	return b
}
```

---

## 11. Maps (`map[K]V`)

### Créer

```go
m := map[string]int{"A": 1}
```

### Lire

```go
v := m["A"]
```

### Tester l'existence d'une clé

```go
v, ok := m["B"]
if !ok {
	fmt.Println("clé absente")
}
```

### Écrire / modifier

```go
m["B"] = 2
```

### Supprimer

```go
delete(m, "A")
```

Piège fréquent:
- Confondre map et slice (clé string vs index int).

---

## 12. Structs

Une struct regroupe des champs.

```go
type Car struct {
	speed   int
	battery int
}
```

Créer une instance:

```go
c := Car{speed: 5, battery: 100}
```

Constructeur idiomatique:

```go
func NewCar(speed int) Car {
	return Car{speed: speed, battery: 100}
}
```

---

## 13. Méthodes

Méthode = fonction attachée à un type.

```go
func (c Car) Speed() int {
	return c.speed
}
```

- Receiver valeur (`c Car`) = copie.

Méthode qui modifie l'objet:

```go
func (c *Car) Drive() {
	if c.battery > 0 {
		c.battery--
	}
}
```

- Receiver pointeur (`*Car`) = modifie l'original.

---

## 14. Pointeurs (`&`, `*`)

### `&x`

Donne l'adresse de `x`.

### `*p`

Accède/modifie la valeur pointée.

```go
x := 10
p := &x
*p = 20
fmt.Println(x) // 20
```

Quand utiliser:
- Modifier une valeur depuis une fonction.
- Éviter de copier de grosses structs.

---

## 15. Interfaces

Une interface décrit des méthodes attendues.

```go
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}
```

Un type implémente l'interface automatiquement s'il a ces méthodes.

```go
type French struct{}

func (French) LanguageName() string { return "French" }
func (French) Greet(name string) string { return "Salut " + name }
```

Pas de mot-clé `implements` en Go.

---

## 16. `any`, assertion de type, type switch

### `any`

Alias de `interface{}` (peut contenir n'importe quel type).

### Assertion de type

```go
v, ok := x.(int)
if ok {
	fmt.Println(v)
}
```

### Type switch

```go
switch v := x.(type) {
case int:
	fmt.Println("int", v)
case string:
	fmt.Println("string", v)
default:
	fmt.Println("autre type")
}
```

---

## 17. Erreurs en Go

### `error`

Une erreur est une valeur.

### Créer une erreur simple

```go
err := errors.New("invalid number of cows")
```

### Pattern standard

```go
value, err := Compute()
if err != nil {
	return 0, err
}
return value, nil
```

### Erreur personnalisée

```go
type InvalidCowsError struct {
	Cows int
	Msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.Cows, e.Msg)
}
```

### Sentinelle

```go
var ErrNotFound = errors.New("not found")
```

---

## 18. Dates et heures (`time`)

### Maintenant

```go
now := time.Now()
```

### Parsing

```go
layout := "2006-01-02 15:04"
t, err := time.Parse(layout, "2026-02-22 14:30")
```

### Comparer

```go
if t.Before(now) {
	fmt.Println("dans le passé")
}
```

### Formater

```go
s := t.Format("Monday, January 2, 2006 at 15:04")
```

Piège fréquent:
- Le layout doit toujours utiliser la date de référence Go.

---

## 19. Documentation Go

### Commentaire de package

```go
// Package weather fournit des utilitaires météo.
package weather
```

### Commentaire d'élément exporté

```go
// Forecast returns a weather forecast string.
func Forecast(city string) string { ... }
```

Règle simple:
- Le commentaire commence par le nom exact de ce qu'il documente.

---

## 20. Recettes rapides (copier/coller)

### Compter des éléments dans une slice

```go
count := 0
for _, v := range items {
	if v == "noodles" {
		count++
	}
}
```

### Récupérer un élément d'index avec garde

```go
func GetItem(s []int, index int) int {
	if index < 0 || index >= len(s) {
		return -1
	}
	return s[index]
}
```

### Supprimer un élément d'index

```go
func RemoveItem(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
```

### Remplacer le dernier élément d'une slice

```go
func AddSecretIngredient(friend []string, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}
```

### Vérifier si une fonction a échoué

```go
result, err := Work()
if err != nil {
	return err
}
_ = result
```

### Type switch standard

```go
func DescribeAnything(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return "string: " + v
	default:
		return "unknown"
	}
}
```

---

## 21. Pièges ultra fréquents (et correction)

1. Oublier `range` dans une boucle
- Faux: `for _, v := s {}`
- Vrai: `for _, v := range s {}`

2. Oublier de réassigner `append`
- Faux: `append(s, 1)`
- Vrai: `s = append(s, 1)`

3. Mauvais ordre des paramètres attendu par les tests
- Vérifier la signature exacte du README/tests.

4. Sortir trop tôt d'une boucle (`default` dans un `switch` interne)
- Mettre le `return default` après la boucle si on doit scanner toute la string.

5. Compter des caractères avec `len` au lieu des runes
- Utiliser `utf8.RuneCountInString`.

6. Utiliser `fmt.Sprint` avec `%d`
- Utiliser `fmt.Sprintf`.

7. Erreur ignorée
- Toujours tester `if err != nil`.

8. Index out of range
- Toujours vérifier: `index < 0 || index >= len(slice)`.

9. Modifier une copie au lieu de l'original
- Si une fonction retourne un nouvel état, réassigner (`x = Update(x)`).
- Si méthode mutante, receiver pointeur (`*T`).

10. Import non utilisé
- Supprimer l'import inutile.

---

## Sources couvertes

- airport-robot
- annalyns-infiltration
- bird-watcher
- blackjack
- booking-up-for-beauty
- card-tricks
- cars-assemble
- census
- chessboard
- election-day
- gross-store
- hello-world
- interest-is-interesting
- jedliks-toys
- lasagna
- lasagna-master
- logs-logs-logs
- meteorology
- need-for-speed
- party-robot
- sorting-room
- the-farm
- vehicle-purchase
- weather-forecast
- welcome-to-tech-palace

---

## 22. Lecture ligne par ligne (explication détaillée)

Cette section prend des extraits fréquents et explique chaque ligne.

### 22.1 Fonction simple avec condition

```go
func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}
	return len(layers) * avg
}
```

Explication ligne par ligne:
- `func PreparationTime(layers []string, avg int) int {`
  - Déclare une fonction nommée `PreparationTime`.
  - Paramètre 1: `layers` (slice de chaînes).
  - Paramètre 2: `avg` (entier).
  - Retourne un `int`.
- `if avg == 0 {`
  - Vérifie si `avg` vaut 0.
  - Ici, 0 sert de signal "valeur par défaut".
- `avg = 2`
  - Remplace `avg` par 2 (valeur par défaut).
- `}`
  - Fin du bloc `if`.
- `return len(layers) * avg`
  - `len(layers)` donne le nombre de couches.
  - Multiplie par le temps moyen `avg`.
  - Retourne le total.
- `}`
  - Fin de la fonction.

### 22.2 Compter des éléments avec `range`

```go
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}
```

Explication ligne par ligne:
- `func Quantities(layers []string) (int, float64) {`
  - Fonction qui retourne 2 valeurs: quantité de nouilles (`int`) et de sauce (`float64`).
- `noodles := 0`
  - Initialise le compteur de grammes de nouilles.
- `sauce := 0.0`
  - Initialise la quantité de sauce en litres.
- `for _, layer := range layers {`
  - Parcourt chaque élément de `layers`.
  - `_` ignore l'index.
  - `layer` reçoit la valeur courante.
- `if layer == "noodles" {`
  - Teste si la couche actuelle est `noodles`.
- `noodles += 50`
  - Ajoute 50g de nouilles.
- `if layer == "sauce" {`
  - Teste si la couche actuelle est `sauce`.
- `sauce += 0.2`
  - Ajoute 0.2L de sauce.
- `return noodles, sauce`
  - Retourne les 2 résultats.

### 22.3 Supprimer un élément d'une slice

```go
func RemoveItem(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
```

Explication ligne par ligne:
- `func RemoveItem(s []int, index int) []int {`
  - Fonction qui prend une slice + un index.
  - Retourne une nouvelle slice (ou la même) sans l'élément ciblé.
- `if index < 0 || index >= len(s) {`
  - Vérifie que l'index est valide.
  - `index < 0` -> index négatif interdit.
  - `index >= len(s)` -> hors limites.
- `return s`
  - Si index invalide, on ne change rien.
- `return append(s[:index], s[index+1:]...)`
  - `s[:index]` = partie avant l'élément.
  - `s[index+1:]` = partie après l'élément.
  - `append(..., ... )` recolle les deux.
  - `...` déplie la 2e slice en arguments variadiques.

### 22.4 Méthode avec receiver pointeur

```go
type Car struct {
	battery int
	speed   int
}

func (c *Car) Drive() {
	if c.battery <= 0 {
		return
	}
	c.battery--
}
```

Explication ligne par ligne:
- `type Car struct { ... }`
  - Définit une structure `Car` avec deux champs.
- `func (c *Car) Drive() {`
  - Méthode attachée à `Car`.
  - `*Car` = pointeur vers la voiture réelle.
  - Donc la méthode peut modifier l'objet original.
- `if c.battery <= 0 { return }`
  - Si batterie vide, on sort immédiatement.
- `c.battery--`
  - Sinon, on retire 1 unité.

### 22.5 Interface et polymorphisme

```go
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
```

Explication ligne par ligne:
- `type Greeter interface { ... }`
  - Contrat: tout type qui a ces 2 méthodes est un `Greeter`.
- `func SayHello(name string, g Greeter) string {`
  - La fonction accepte n'importe quel type implémentant `Greeter`.
- `g.LanguageName()`
  - Appelle la méthode du type concret passé.
- `g.Greet(name)`
  - Appelle la salutation du type concret.
- `fmt.Sprintf(...)`
  - Construit la string finale.

### 22.6 Gestion d'erreur standard

```go
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}
```

Explication ligne par ligne:
- `func ... (float64, error)`
  - Retourne une valeur utile + une erreur éventuelle.
- `if cows <= 0 {`
  - Vérifie l'entrée avant d'appeler d'autres fonctions.
- `return 0, errors.New("invalid number of cows")`
  - En cas d'erreur, valeur utile à zéro + message d'erreur.
- `return DivideFood(fc, cows)`
  - Cas valide: délègue au calcul principal et renvoie ses résultats.

### 22.7 Type assertion et type switch

```go
func DescribeAnything(i any) string {
	switch v := i.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return "string: " + v
	default:
		return "unknown"
	}
}
```

Explication ligne par ligne:
- `i any`
  - Paramètre pouvant contenir n'importe quel type.
- `switch v := i.(type) {`
  - Détecte dynamiquement le type réel stocké dans `i`.
- `case int:`
  - Si c'est un entier, `v` est de type `int` dans ce bloc.
- `case string:`
  - Si c'est une string, `v` est de type `string`.
- `default:`
  - Tous les autres types.

### 22.8 String, runes, UTF-8

```go
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
```

Explication ligne par ligne:
- `log string`
  - Texte potentiellement avec accents/emojis.
- `utf8.RuneCountInString(log)`
  - Compte les runes (caractères Unicode), pas les octets.
- `<= limit`
  - Renvoie `true` si le texte respecte la limite.

### 22.9 Dates avec `time.Parse`

```go
layout := "2006-01-02 15:04"
t, err := time.Parse(layout, "2026-02-22 14:30")
if err != nil {
	return err
}
```

Explication ligne par ligne:
- `layout := "2006-01-02 15:04"`
  - Format attendu de la date texte.
  - En Go, on utilise la date de référence `2006-01-02 15:04:05`.
- `time.Parse(layout, "...")`
  - Convertit une string en `time.Time`.
- `if err != nil { ... }`
  - Obligatoire pour gérer un format invalide.

### 22.10 Lecture d'une fonction testée (mental model)

Quand tu lis une fonction, pose-toi ces questions ligne par ligne:
- Quels sont les types d'entrée ?
- Y a-t-il une garde de sécurité au début (`if invalid -> return`) ?
- La boucle parcourt quoi exactement (`index`, `value`, `key`, `rune`) ?
- Où l'état change (`count++`, `append`, `field = ...`) ?
- Le retour final correspond bien au type attendu par les tests ?

