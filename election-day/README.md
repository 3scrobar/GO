# Jour d'élection

Bienvenue dans « Jour d'élection » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Comme beaucoup d’autres langages, Go possède des pointeurs.
Si vous débutez avec les pointeurs, ils peuvent sembler un peu mystérieux, mais une fois que vous vous y êtes habitué, ils sont assez simples.
Ils constituent une partie cruciale du Go, alors prenez le temps de bien les comprendre.

Avant d’entrer dans les détails, il convient de comprendre l’utilisation des pointeurs. Les pointeurs sont un moyen de partager de la mémoire avec d'autres parties de notre programme, ce qui est utile pour deux raisons principales :
1. Lorsque nous disposons de grandes quantités de données, faire des copies à transmettre entre les fonctions est très inefficace.
En transmettant l'emplacement mémoire où les données sont stockées, nous pouvons réduire considérablement l'empreinte des ressources de nos programmes.
2. En passant des pointeurs entre les fonctions, nous pouvons accéder et modifier directement la copie unique des données, ce qui signifie que toutes les modifications apportées par une fonction sont immédiatement visibles par les autres parties du programme lorsque la fonction se termine.

## Variables et mémoire

Disons que nous avons une variable entière régulière `a` :

```go
var a int
```

Lorsque l'on déclare une variable, Go doit trouver une place en mémoire pour stocker sa valeur. Ceci nous est largement abstrait : lorsque nous avons besoin de récupérer la valeur stockée dans cet élément de mémoire, nous pouvons simplement y faire référence par le nom de la variable.

Par exemple, lorsque nous écrivons `a + 2`, nous récupérons effectivement la valeur stockée dans la mémoire associée à la variable `a` et y ajoutons 2.

De même, lorsque nous devons modifier la valeur dans la partie mémoire de `a`, nous pouvons utiliser le nom de la variable pour effectuer une affectation :

```go
a = 3
```

La partie de mémoire associée à `a` stockera désormais la valeur `3`.

## Pointeurs

Bien que les variables nous permettent de faire référence à des valeurs en mémoire, il est parfois utile de connaître l'**adresse mémoire** vers laquelle pointe la variable. Les **pointeurs** contiennent les adresses mémoire de ces valeurs. Vous déclarez une variable avec un type pointeur en préfixant le type sous-jacent avec un astérisque :

```go
var p *int // 'p' contains the memory address of an integer
```

Ici, nous déclarons une variable `p` de type "pointeur vers int" (`*int`). Cela signifie que `p` contiendra l'adresse mémoire d'un entier. La valeur zéro des pointeurs est `nil` car un pointeur `nil` ne contient aucune adresse mémoire.

### Obtenir un pointeur vers une variable

Pour trouver l'adresse mémoire de la valeur d'une variable, on peut utiliser l'opérateur `&`.
Par exemple, si nous voulons rechercher et stocker l'adresse mémoire de la variable `a` dans le pointeur `p`, nous pouvons procéder comme suit :

```go
var a int
a = 2

var p *int
p = &a // the variable 'p' contains the memory address of 'a'
```

### Accéder à la valeur via un pointeur (déréférencement)

Lorsque nous avons un pointeur, nous pouvons vouloir connaître la valeur stockée dans l’adresse mémoire représentée par le pointeur. Nous pouvons le faire en utilisant l'opérateur `*` :

```go
var a int
a = 2

var p *int
p = &a // the variable 'p' contains the memory address of 'a'

var b int
b = *p // b == 2
```

L'opération `*p` récupère la valeur stockée à l'adresse mémoire stockée dans `p`. Cette opération est souvent appelée « déréférencement ».

On peut également utiliser l'opérateur de déréférencement pour attribuer une nouvelle valeur à l'adresse mémoire référencée par le pointeur :

```go
var a int        // declare int variable 'a'
a = 2            // assign 'a' the value of 2

var pa *int
pa = &a          // 'pa' now contains to the memory address of 'a'
*pa = *pa + 2    // increment by 2 the value at memory address 'pa'

fmt.Println(a)   // Output: 4
                 // 'a' will have the new value that was changed via the pointer!
```

L'affectation à `*pa` modifiera la valeur stockée à l'adresse mémoire `pa`. Puisque `pa` contient l'adresse mémoire de `a`, en l'attribuant à `*pa`, nous modifions effectivement la valeur de `a` !

Attention toutefois : vérifiez toujours si un pointeur n'est pas `nil` avant de déréférencer. Le déréférencement d'un pointeur `nil` fera planter le programme au moment de l'exécution !

```go
var p *int // p is nil initially
fmt.Println(*p)
// panic: runtime error: invalid memory address or nil pointer dereference
```

### Pointeurs vers des structures

Jusqu'à présent, nous n'avons vu que des pointeurs vers des valeurs primitives. Nous pouvons également créer des pointeurs pour les structures :

```go
type Person struct {
    Name string
    Age  int
}

var peter Person
peter = Person{Name: "Peter", Age: 22}

var p *Person
p = &peter
```

Nous aurions également pu créer un nouveau `Person` et y stocker immédiatement un pointeur :

```go
var p *Person
p = &Person{Name: "Peter", Age: 22}
```

Lorsque nous avons un pointeur vers une structure, nous n'avons pas besoin de déréférencer le pointeur avant d'accéder à l'un des champs :

```go
var p *Person
p = &Person{Name: "Peter", Age: 22}

fmt.Println(p.Name) // Output: "Peter"
                    // Go automatically dereferences 'p' to allow
                    // access to the 'Name' field
```

## Les tranches et les cartes sont déjà des pointeurs

Les tranches et les cartes sont des types spéciaux car elles comportent déjà des pointeurs dans leur implémentation. Cela signifie que le plus souvent, nous n'avons pas besoin de créer des pointeurs pour ces types afin de partager l'adresse mémoire de leurs valeurs. Imaginez que nous ayons une fonction qui incrémente la valeur d'une clé dans une carte :


```go
func incrementPeterAge(m map[string]int) {
	m["Peter"] += 1
}
```

Si nous créons une carte et appelons cette fonction, les modifications apportées par la fonction à la carte persistent après la fin de la fonction. Il s'agit d'un comportement similaire que nous obtenons si nous utilisons un pointeur, mais notez que dans cet exemple, nous n'utilisons aucun référencement/déréférencement ni aucune syntaxe de pointeur :

```go
ages := map[string]int{
  "Peter": 21
}
incrementPeterAge(ages)
fmt.Println(ages)
// Output: map[Peter:22]
// The changes the function 'incrementPeterAge' made to the map are visible after the function ends!
```

La même chose s'applique lors de la modification d'un élément existant dans une tranche.

Cependant, les actions qui renvoient une nouvelle tranche comme `append` constituent un cas particulier et **pourraient ne pas** modifier la tranche en dehors de la fonction. Cela est dû au fonctionnement interne des tranches, mais nous n’aborderons pas cela en détail dans cet exercice, car il s’agit d’un sujet plus avancé. Si vous êtes vraiment curieux, vous pouvez en savoir plus à ce sujet dans [Go Blog : Mécanique de 'append'][mechanics-of-append]

[mechanics-of-append]: https://go.dev/blog/slices

## Instructions

Une école locale près de chez vous possède une association d'étudiants très active.
L'association des étudiants est dirigée par un président et une fois tous les 2 ans,
des élections sont organisées pour élire un nouveau président.

Lors des élections de cette année, il a été décidé qu'un nouveau système numérique
il fallait compter les voix. L'école a besoin de votre aide pour construire ce nouveau système.

## 1. Créez un compteur de votes

L’une des premières choses dont le nouveau système de vote a besoin est un compteur de votes.
Ce compteur est un moyen de suivre les votes d'un candidat particulier.

Créez une fonction `NewVoteCounter` qui accepte le nombre de votes initiaux pour un candidat et renvoie un pointeur faisant référence à un `int`, initialisé avec le nombre de votes initiaux donné.

```go
var initialVotes int
initialVotes = 2

var counter *int
counter = NewVoteCounter(initialVotes)
*counter == initialVotes // true
```

## 2. Obtenez le nombre de votes d'un compteur

Vous avez maintenant un moyen de créer de nouveaux compteurs ! Mais vous réalisez maintenant que le nouveau système aura également besoin d’un moyen d’obtenir le nombre de voix à partir d’un compteur.

Créez une fonction `VoteCount` qui prendra un compteur (`*int`) comme argument et renverra le nombre de votes dans le compteur. Si le compteur est `nil`, vous devez supposer qu'il n'a aucun vote :

```go
var votes int
votes = 3

var voteCounter *int
voteCounter = &votes

VoteCount(voteCounter)
// => 3

var nilVoteCounter *int
VoteCount(nilVoteCounter)
// => 0
```

## 3. Augmenter les votes d'un compteur

C'est enfin le moment de compter les votes ! Vous avez maintenant besoin d'un moyen d'augmenter les votes dans un compteur.

Créez une fonction `IncrementVoteCount` qui prendra un compteur (`*int`) comme argument et un nombre de votes, et incrémentera le compteur de ce nombre de votes. Vous pouvez supposer que le pointeur transmis ne sera jamais `nil`.

```go
var votes int
votes = 3

var voteCounter *int
voteCounter = &votes

IncrementVoteCount(voteCounter, 2)

votes == 5          // true
*voteCounter == 5   // true
```

## 4. Créer les résultats des élections

Tous les votes étant désormais comptés, il est temps de préparer l'annonce des résultats à toute l'école.
Pour cela, vous constatez qu'avoir uniquement des compteurs pour les votes est insuffisant.
Il doit y avoir un moyen d'associer le nombre de voix à un candidat particulier.

Créez une fonction `NewElectionResult` qui reçoit le nom d'un candidat et son nombre de voix et
renvoie un nouveau résultat d'élection.

```go
var result *ElectionResult
result = NewElectionResult("Peter", 3)

result.Name == "Peter"  // true
result.Votes == 3       // true
```

La structure des résultats des élections est déjà créée pour vous et est définie comme :

```go
type ElectionResult struct {
    // Name of the candidate
    Name    string
    // Votes of votes the candidate had
    Votes   int
}
```

## 5. Annoncer les résultats

Il est temps d'annoncer le nouveau président de l'école !
Le président sera annoncé sur les petits babillards numériques dont dispose l'école.
Le message doit indiquer le nom du nouveau président et les voix qu'il a obtenues, au format suivant : `<candidate_name> (<votes>)`. Voici un exemple d'un tel message : `"Peter (51)"`.

Créez une fonction `DisplayResult` qui recevra un `*ElectionResult` en argument et renverra une chaîne avec le message à afficher.


```go
var result *ElectionResult
result = &ElectionResult{
    Name: "John",
    Votes: 32,
}

DisplayResult(result)
// => John (32)
```

## 6. Recomptage des votes

Pour garantir l'exactitude des résultats finaux, les votes ont été recomptés. Lors du recomptage, il a été constaté que le nombre de voix pour certains candidats était inférieur d'une unité.

Créez une fonction `DecrementVotesOfCandidate` qui reçoit les résultats finaux et le nom d'un candidat pour lequel vous devez décrémenter son décompte des voix. Les résultats finaux sont donnés sous la forme d'un `map[string]int`, où les clés sont les noms des candidats et les valeurs sont le total de leurs voix.

```go
var finalResults = map[string]int{
    "Mary":  10,
    "John":  51,
}

DecrementVotesOfCandidate(finalResults, "Mary")

finalResults["Mary"]
// => 9
```

## Source

### Créé par

- @andrerfcsantos
