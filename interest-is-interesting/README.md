# L'intérêt est intéressant

Bienvenue dans « L'intérêt est intéressant » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Un nombre à virgule flottante est un nombre comportant zéro ou plusieurs chiffres derrière le séparateur décimal. Exemples : `-2.4`, `0.1`, `3.14`, `16.984025` et `1024.0`.

Différents types à virgule flottante peuvent stocker différents nombres de chiffres après le séparateur de chiffres : c'est ce qu'on appelle leur précision.

Go a deux types à virgule flottante :

- `float32` : 32 bits (précision de ~6-9 chiffres).
- `float64` : 64 bits (précision de ~15-17 chiffres). Il s'agit du type à virgule flottante par défaut.

Comme on peut le constater, les deux types peuvent stocker un nombre différent de chiffres. Cela signifie qu'essayer de stocker PI dans un `float32` ne stockera que les 6 à 9 premiers chiffres (le dernier chiffre étant arrondi).

Par défaut, Go utilisera `float64` pour les nombres à virgule flottante, sauf si le nombre à virgule flottante est :

1. affecté à une variable de type `float32`, ou
2. renvoyé par une fonction avec le type de retour `float32`, ou
3. passé en argument à la fonction `float32()`.

Le package `math` contient de nombreuses fonctions mathématiques utiles.

## Instructions

Dans cet exercice, vous travaillerez avec des comptes d'épargne.
Chaque année, le solde de votre compte épargne est mis à jour en fonction de son taux d'intérêt.
Le taux d'intérêt que votre banque vous accorde dépend du montant d'argent sur votre compte (son solde) :

- 3,213% pour un solde inférieur à `0` dollars (le solde devient encore plus négatif).
- 0,5% pour un solde supérieur ou égal à `0` dollars, et inférieur à `1000` dollars.
- 1,621% pour un solde supérieur ou égal à `1000` dollars, et inférieur à `5000` dollars.
- 2,475% pour un solde supérieur ou égal à `5000` dollars.

Vous avez quatre tâches, dont chacune traitera de votre solde et de son taux d'intérêt.

## 1. Calculez le taux d'intérêt

Implémentez la fonction `InterestRate()` pour calculer le taux d'intérêt en fonction du solde spécifié :

```go
InterestRate(200.75)
// => 0.5
```

Notez que la valeur renvoyée est un `float32`.

## 2. Calculez les intérêts

Implémentez la fonction `Interest()` pour calculer les intérêts en fonction du solde spécifié :

```go
Interest(200.75)
// => 1.003750
```

Notez que la valeur renvoyée est un `float64`.

## 3. Calculer la mise à jour du solde annuel

Implémentez la fonction `AnnualBalanceUpdate()` pour calculer la mise à jour annuelle du solde, en tenant compte du taux d'intérêt :

```go
AnnualBalanceUpdate(200.75)
// => 201.75375
```

Notez que la valeur renvoyée est un `float64`.

## 4. Calculez les années avant d'atteindre l'équilibre souhaité

Implémentez la fonction `YearsBeforeDesiredBalance()` pour calculer le nombre minimum d'années nécessaires pour atteindre le solde souhaité, en tenant compte du fait que chaque année, des intérêts sont ajoutés au solde.
Cela signifie que le solde après un an est le suivant : solde de départ + intérêts pour le solde de départ.
Le solde après la deuxième année est : solde après un an + intérêts pour solde après un an.
Et ainsi de suite, jusqu'à ce que le solde de l'année en cours soit supérieur ou égal au solde cible.

```go
balance := 200.75
targetBalance := 214.88
YearsBeforeDesiredBalance(balance, targetBalance)
// => 14
```

Notez que la valeur renvoyée est un `int`.

## Source

### Créé par

- @erikschierboom
