# Référence Go - Version Débutant (Exercism)

Ce document est un mémo Go orienté débutant.
Il regroupe les notions des README Exercism et explique chaque commande/outil de base avec des exemples simples.

## Comment utiliser ce document

- Si tu bloques sur une ligne de code, va d'abord à l'`Index rapide`.
- Lis ensuite la section complète avec l'exemple.
- Reviens coder, puis valide avec les tests.
- Ordre conseillé pour débuter: sections 1 -> 10, puis 17, puis 22.

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
- Lecture ligne par ligne (explication guidée) -> section 22
- Aléatoire (`math/rand`: `Intn`, `Float64`, `Shuffle`) -> section 23
- Expressions régulières (`regexp`: `MustCompile`, `MatchString`, `Split`, `ReplaceAllString`) -> section 24
- Fonctions d'ordre supérieur / closures (prédicats) -> section 25
- Matrice exercice -> notions (audit) -> section 26

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
- animal-magic
- annalyns-infiltration
- bird-watcher
- blackjack
- booking-up-for-beauty
- card-tricks
- cars-assemble
- census
- chessboard
- election-day
- expenses
- gross-store
- hello-world
- interest-is-interesting
- jedliks-toys
- lasagna
- lasagna-master
- logs-logs-logs
- meteorology
- need-for-speed
- parsing-log-files
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



### 22.11 Validation d'une ligne de log avec regex

```go
var validLineRE = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

func IsValidLine(text string) bool {
	return validLineRE.MatchString(text)
}
```

Explication ligne par ligne:
- `var validLineRE = regexp.MustCompile(...)`
  - Compile la regex une fois au chargement du package.
  - `^` = début de ligne.
  - `\[` et `\]` = crochets littéraux.
  - `(TRC|DBG|...)` = une des options autorisées.
- `func IsValidLine(text string) bool {`
  - Fonction qui prend une ligne de log et renvoie vrai/faux.
- `return validLineRE.MatchString(text)`
  - Vérifie si `text` correspond au motif au début de la ligne.

### 22.12 Découper une ligne avec un séparateur regex

```go
var splitLogLineRE = regexp.MustCompile(`<[~*=-]*>`)

func SplitLogLine(text string) []string {
	return splitLogLineRE.Split(text, -1)
}
```

Explication ligne par ligne:
- `<[~*=-]*>`
  - Début `<`, fin `>`.
  - Entre les deux: zéro ou plusieurs caractères parmi `~`, `*`, `=`, `-`.
  - `*` permet aussi `<>`.
- `Split(text, -1)`
  - Coupe `text` partout où le séparateur est trouvé.
  - `-1` = pas de limite sur le nombre de morceaux.

### 22.13 Tag utilisateur avec capture de groupe

```go
var userNameCaptureRE = regexp.MustCompile(`User\s+([^\s]+)\s`)

func TagWithUserName(lines []string) []string {
	out := make([]string, len(lines))
	for i, line := range lines {
		m := userNameCaptureRE.FindStringSubmatch(line)
		if len(m) == 2 {
			out[i] = "[USR] " + m[1] + " " + line
		} else {
			out[i] = line
		}
	}
	return out
}
```

Explication ligne par ligne:
- `User\s+`
  - Cherche le mot `User` suivi d'un ou plusieurs espaces.
- `([^\s]+)`
  - Groupe capturé: le nom d'utilisateur (suite sans espaces).
- `\s`
  - Un espace après le nom (hypothèse du sujet).
- `out := make([]string, len(lines))`
  - Prépare une slice résultat de même taille.
- `FindStringSubmatch(line)`
  - Renvoie la correspondance complète + les groupes capturés.
- `len(m) == 2`
  - `m[0]` = match complet, `m[1]` = nom capturé.

### 22.14 Calcul avec `rand` ligne par ligne

```go
func RollADie() int {
	return rand.Intn(20) + 1
}

func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}
```

Explication ligne par ligne:
- `rand.Intn(20)`
  - Donne un entier dans `[0, 20)`.
- `+ 1`
  - Décale vers `[1, 20]`.
- `rand.Float64()`
  - Donne un float dans `[0.0, 1.0)`.
- `* 12.0`
  - Échelle vers `[0.0, 12.0)`.

### 22.15 Closure de filtre (prédicat)

```go
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}
```

Explication ligne par ligne:
- `ByCategory(c string)`
  - Prend une catégorie cible (`c`).
- `func(Record) bool`
  - Retourne une fonction prédicat.
- `return func(r Record) bool { ... }`
  - Fonction anonyme qui "capture" `c`.
- `r.Category == c`
  - Vrai uniquement pour les enregistrements de la catégorie demandée.

### 22.16 Erreur personnalisée pas à pas

```go
type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}
```

Explication ligne par ligne:
- `type InvalidCowsError struct { ... }`
  - Type dédié pour transporter le contexte de l'erreur.
- `func (e *InvalidCowsError) Error() string`
  - Implémente l'interface `error`.
- `Sprintf(...)`
  - Formate un message utile pour debug et tests.

### 22.17 Itération map ligne par ligne

```go
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	for _, file := range cb {
		if len(file) >= rank && file[rank-1] {
			count++
		}
	}
	return count
}
```

Explication ligne par ligne:
- `if rank < 1 || rank > 8`
  - Garde de sécurité pour éviter un index invalide.
- `for _, file := range cb`
  - Parcourt chaque colonne (map de `File`).
- `file[rank-1]`
  - Convertit un rang 1..8 en index 0..7.
- `count++`
  - Compte uniquement les cases occupées (`true`).

---

## 23. Aléatoire avec `math/rand`

Notions vues dans `animal-magic`.

### Fonctions principales
- `rand.Intn(n)` -> entier aléatoire dans `[0, n)`.
- `rand.Float64()` -> float aléatoire dans `[0.0, 1.0)`.
- `rand.Shuffle(n, swapFn)` -> mélange des éléments en place.

### Exemples
```go
// 1..20 inclus
roll := rand.Intn(20) + 1

// 0.0 <= e < 12.0
energy := rand.Float64() * 12.0

animals := []string{"ant", "beaver", "cat", "dog", "eel", "fox", "goat", "hedgehog"}
rand.Shuffle(len(animals), func(i, j int) {
	animals[i], animals[j] = animals[j], animals[i]
})
```

### Pièges fréquents
- Oublier que `Intn(20)` commence à 0 (pas 1).
- Oublier que `Float64()` ne prend pas d'argument.
- Oublier de retourner la slice après `Shuffle`.

---

## 24. Expressions régulières (`regexp`)

Notions vues dans `parsing-log-files`.

### Pourquoi compiler une regex
- `regexp.MustCompile(...)` compile une fois et réutilise le motif.
- À placer en variable package si utilisé souvent.

### Méthodes courantes
- `re.MatchString(s)` -> booléen (correspondance oui/non)
- `re.Split(s, -1)` -> découpe la chaîne selon le séparateur regex
- `re.ReplaceAllString(s, repl)` -> remplace toutes les occurrences
- `re.FindStringSubmatch(s)` -> récupère la capture complète + groupes

### Exemples
```go
validLineRE := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
valid := validLineRE.MatchString("[INF] System ready")

sepRE := regexp.MustCompile(`<[~*=-]*>`)
parts := sepRE.Split("section 1<*>section 2<>section 3", -1)

eolRE := regexp.MustCompile(`end-of-line\d+`)
clean := eolRE.ReplaceAllString("[INF] end-of-line23033 ok", "")

userRE := regexp.MustCompile(`User\s+([^\s]+)\s`)
m := userRE.FindStringSubmatch("[WRN] User James123 lost connection.")
// m[1] == "James123"
```

### Pièges fréquents
- `+` veut dire 1 ou plus; pour accepter `<>`, il faut `*` (0 ou plus).
- Oublier l'ancre `^` quand on veut "commence par".
- Compter des occurrences avec `MatchString` alors qu'on veut toutes les occurrences (utiliser `FindAll...`).

---

## 25. Fonctions d'ordre supérieur et closures (prédicats)

Notions vues dans `expenses`.

### Idée
Une fonction peut:
- être passée en argument,
- être renvoyée par une autre fonction,
- capturer des variables (closure).

### Exemple de filtre générique
```go
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := []Record{}
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}
```

### Closure par période
```go
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}
```

### Closure par catégorie
```go
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}
```

### Composition utile
```go
inPeriod := Filter(records, ByDaysPeriod(period))
inCategory := Filter(inPeriod, ByCategory("food"))
```

### Pièges fréquents
- Retourner un bool au lieu de retourner une fonction.
- Oublier que la closure capture les variables extérieures.
- Multiplier les boucles quand un `Filter` composable suffit.


---

## 26. Matrice exercice -> notions (audit)

Cette matrice est extraite automatiquement des README de chaque dossier.
Elle sert de checklist pour retrouver rapidement où une notion a été vue.

- `airport-robot`: Interface comme ensemble de méthodes; Implémenter une interface; Vidéo de l'interface
- `animal-magic`: Graines
- `annalyns-infiltration`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `bird-watcher`: Syntaxe générale; Boucles For - Un exemple
- `blackjack`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `booking-up-for-beauty`: Options de mise en page
- `card-tricks`: Tranches (Tranches); Fonctions Variadiques
- `cars-assemble`: Numéros; Opérateurs arithmétiques; Conversion entre les types; Opérations arithmétiques sur différents types
- `census`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `chessboard`: Itérer sur une slice; Itérer sur une map; Itération en omettant la clé ou la valeur; Types non structurés
- `election-day`: Variables et mémoire; Pointeurs; Obtenir un pointeur vers une variable; Accéder à la valeur via un pointeur (déréférencement); Pointeurs vers des structures; Les tranches et les cartes sont déjà des pointeurs
- `expenses`: Types de fonction; Fonctions anonymes
- `gross-store`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `hello-world`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `interest-is-interesting`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `jedliks-toys`: (pas de sous-section de notion explicite entre Introduction et Instructions)
- `lasagna`: Forfaits; Variables; Constantes; Fonctions
- `lasagna-master`: Paramètres de fonction; Paramètres vs. Arguments; Valeurs de retour; Appel de fonctions; Valeurs de retour nommées et retour nu; Passage par valeur vs. passage par référence; Pointeurs; Exceptions
- `logs-logs-logs`: Unicode et points de code Unicode; UTF-8; Utilisation des Runes; Runes et Chaînes
- `meteorology`: Exemple : Distances
- `need-for-speed`: Définir une struct; Créer des instances d'une struct; Fonctions "Nouveau"
- `parsing-log-files`: Syntaxe; Compilation des modèles - type `RegExp`; Méthodes d'expression régulière; Exemples de `MatchString`; Exemples de `FindString`; Exemples de `FindStringSubmatch`; Exemples de `ReplaceAllString`; Exemples de `Split`
- `party-robot`: Forfaits; Formatage des Chaînes
- `sorting-room`: Conversion de types; Conversion entre types primitifs et chaînes de caractères; Assertions de type; Commutateurs de Type
- `the-farm`: L'interface d'erreur; Créer et renvoyer une erreur simple; Vérification des erreurs; Types d'erreurs personnalisés
- `vehicle-purchase`: Comparaison; Si les déclarations
- `weather-forecast`: Commentaires de documentation; Commentaires de paquet; Commentaires de fonction
- `welcome-to-tech-palace`: (pas de sous-section de notion explicite entre Introduction et Instructions)

---

## 27. Lecture caractère par caractère (niveau débutant absolu)

Objectif: comprendre exactement ce que chaque symbole fait dans une ligne Go.

### 27.1 Regex de validation de log, symbole par symbole

```go
var validLineRE = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
```

Décomposition:
- `var`
  - Déclare une variable.
- `validLineRE`
  - Nom de la variable (ici une regex compilée).
- `=`
  - Affectation.
- `regexp.MustCompile(...)`
  - Compile la regex tout de suite.
  - `MustCompile` panique si la regex est invalide.
- Les backticks `` `...` ``
  - Chaîne brute (pratique pour regex: moins d'échappements Go).
- `^`
  - Début de chaîne.
- `\[` et `\]`
  - Crochets littéraux `[` et `]`.
- `(A|B|C)`
  - Groupe: une option parmi plusieurs.
- `TRC|DBG|INF|WRN|ERR|FTL`
  - Niveaux de logs autorisés.

```go
func IsValidLine(text string) bool {
	return validLineRE.MatchString(text)
}
```

Décomposition:
- `func`
  - Déclare une fonction.
- `IsValidLine`
  - Nom de la fonction.
- `(text string)`
  - Un paramètre nommé `text` de type `string`.
- `bool`
  - Type de retour.
- `return`
  - Renvoie le résultat.
- `MatchString(text)`
  - Teste si `text` matche la regex.

### 27.2 Closure de filtre, symbole par symbole

```go
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}
```

Décomposition:
- `func ByCategory(c string)`
  - Fonction qui reçoit une catégorie cible `c`.
- `func(Record) bool`
  - Type de retour: "une fonction qui prend `Record` et renvoie `bool`".
- `return func(r Record) bool { ... }`
  - Retourne une fonction anonyme.
- `r.Category == c`
  - Compare la catégorie du record à la valeur capturée `c`.
- Capture (closure)
  - La fonction interne "se souvient" de `c` même après la fin de `ByCategory`.

### 27.3 Suppression dans une slice avec `append`, symbole par symbole

```go
return append(s[:index], s[index+1:]...)
```

Décomposition:
- `s[:index]`
  - Prend les éléments avant `index`.
- `s[index+1:]`
  - Prend les éléments après `index`.
- `...`
  - "Déplie" la 2e slice en arguments variadiques pour `append`.
- `append(a, b...)`
  - Concatène `a` et `b`.
- Résultat
  - Même slice sans l'élément d'index `index`.

Version complète avec garde:

```go
func RemoveItem(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}
```

Décomposition de la garde:
- `index < 0`
  - Évite un index négatif.
- `index >= len(s)`
  - Évite un dépassement de slice.
- `||`
  - OU logique: si une des deux conditions est vraie, index invalide.

### 27.4 Lecture d'une signature de fonction (anti-panique)

Exemple:

```go
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error)
```

Comment lire:
- Nom: `CategoryExpenses`
- Entrées:
  - `in` -> `[]Record`
  - `p` -> `DaysPeriod`
  - `c` -> `string`
- Sorties:
  - `float64` (valeur utile)
  - `error` (erreur éventuelle)

Règle pratique:
- Si tu vois `(X, error)`, pense immédiatement:
  - "je dois gérer `err != nil` côté appelant".

### 27.5 Mini-glossaire des symboles (ultra rapide)

- `:=` -> déclare + assigne (dans une fonction)
- `=` -> assigne une variable existante
- `==` -> compare l'égalité
- `!=` -> compare la différence
- `&&` -> ET logique
- `||` -> OU logique
- `[]T` -> slice de type `T`
- `map[K]V` -> map clé `K`, valeur `V`
- `*T` -> pointeur vers `T`
- `&x` -> adresse de `x`
- `...` -> arguments variadiques / dépliage de slice
- `_` -> ignore volontairement une valeur

