# L'infiltration d'Annalyn

Bienvenue dans « L'infiltration d'Annalyn » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Les booléens en Go sont représentés par le type booléen prédéclaré `bool`, dont les valeurs peuvent être soit `true` soit `false`.
C'est un type défini.

```go
var closed bool    // variable booléenne 'closed' initialisée implicitement avec 'false'
speeding := true   // variable booléenne 'speeding' initialisée avec 'true'
hasError := false  // variable booléenne 'hasError' initialisée avec 'false' 
```

Go prend en charge trois opérateurs logiques qui peuvent évaluer des expressions en valeurs booléennes, retournant soit `true` soit `false`.

| Opérateur   | Ce que cela signifie                                            |
| ----------- | -------------------------------------------------------------- |
| `&&` (et)   | C'est vrai si les deux énoncés sont vrais.                    |
| `\|\|` (ou) | C'est vrai si au moins un énoncé est vrai.                   |
| `!` (non)   | C'est vrai uniquement si l'énoncé est faux.                  |

## Instructions

Dans cet exercice, vous implémenterez la logique de quête pour un nouveau jeu RPG qu'un ami développe. Le personnage principal du jeu est Annalyn, une fille courageuse avec un chien de compagnie féroce et loyal. Malheureusement, un désastre frappe : sa meilleure amie a été kidnappée en cherchant des baies dans la forêt. Annalyn essaiera de trouver et de libérer sa meilleure amie, en prenant éventuellement son chien avec elle dans cette quête.

Après avoir suivi la piste de sa meilleure amie pendant un certain temps, elle trouve le camp où son amie est emprisonnée. Il s'avère qu'il y a deux ravisseurs : un chevalier puissant et un archer rusé.

Ayant trouvé les ravisseurs, Annalyn envisage laquelle des actions suivantes elle peut entreprendre :

- _Attaque rapide_ : une attaque rapide peut être faite si le chevalier dort, car il lui faut du temps pour mettre son armure, il sera donc vulnérable.
- _Espionnage_ : le groupe peut être espionné si au moins l'un d'eux est éveillé. Sinon, l'espionnage est une perte de temps.
- _Signaler le prisonnier_ : le prisonnier peut être signalé en utilisant des sons d'oiseaux si le prisonnier est éveillé et l'archer dort, car les archers sont formés à la signalisation par oiseaux pour qu'ils puissent intercepter le message.
- _Libérer le prisonnier_ : Annalyn peut essayer de s'infiltrer discrètement dans le camp pour libérer le prisonnier.
C'est une chose risquée et ne peut réussir que de deux façons :
- Si Annalyn a son chien de compagnie avec elle, elle peut sauver le prisonnier si l'archer dort.
Le chevalier a peur du chien et l'archer n'aura pas le temps de se préparer avant qu'Annalyn et le prisonnier ne s'échappent.
- Si Annalyn n'a pas son chien, elle et le prisonnier doivent être très discrets !
Annalyn peut libérer le prisonnier si le prisonnier est éveillé et que le chevalier et l'archer dorment tous les deux, mais si le prisonnier dort, il ne peut pas être sauvé : le prisonnier serait surpris par l'apparition soudaine d'Annalyn et réveillerait le chevalier et l'archer.

Vous avez quatre tâches : implémenter la logique pour déterminer si les actions ci-dessus sont disponibles en fonction de l'état des trois personnages trouvés dans la forêt et de la présence ou non du chien de compagnie d'Annalyn.

## 1. Vérifier si une attaque rapide peut être faite

Définissez la fonction `CanFastAttack()` qui prend une valeur booléenne indiquant si le chevalier est éveillé. Cette fonction retourne `true` si une attaque rapide peut être faite en fonction de l'état du chevalier. Sinon, elle retourne `false` :

```go
var knightIsAwake = true
fmt.Println(CanFastAttack(knightIsAwake))
// Résultat : false
```

## 2. Vérifier si le groupe peut être espionné

Définissez la fonction `CanSpy()` qui prend trois valeurs booléennes, indiquant respectivement si le chevalier, l'archer et le prisonnier sont éveillés. La fonction retourne `true` si le groupe peut être espionné, en fonction de l'état des trois personnages. Sinon, elle retourne `false` :

```go
var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
// Résultat : true
```

## 3. Vérifier si le prisonnier peut être signalé

Définissez la fonction `CanSignalPrisoner()` qui prend deux valeurs booléennes, indiquant respectivement si l'archer et le prisonnier sont éveillés. La fonction retourne `true` si le prisonnier peut être signalé, en fonction de l'état des deux personnages. Sinon, elle retourne `false` :

```go
var archerIsAwake = false
var prisonerIsAwake = true
fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
// Résultat : true
```

## 4. Vérifier si le prisonnier peut être libéré

Définissez la fonction `CanFreePrisoner()` qui prend quatre valeurs booléennes. Les trois premiers paramètres indiquent respectivement si le chevalier, l'archer et le prisonnier sont éveillés. Le dernier paramètre indique si le chien de compagnie d'Annalyn est présent. La fonction retourne `true` si le prisonnier peut être libéré en fonction de l'état des trois personnages et de la présence du chien de compagnie d'Annalyn. Sinon, elle retourne `false` :

```go
var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false
fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
// Résultat : false
```

## Source

### Créé par

- @oanaOM
