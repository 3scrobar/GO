# L'intérêt, c'est intéressant

Bienvenue sur L'intérêt, c'est intéressant dans la piste Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans les utiliser :)

## Introduction

Un nombre flottant est un nombre avec zéro ou plusieurs chiffres après le séparateur décimal. Les exemples sont `-2.4`, `0.1`, `3.14`, `16.984025` et `1024.0`.

Différents types flottants peuvent stocker un nombre différent de chiffres après le séparateur décimal - c'est ce qu'on appelle sa précision.

Go a deux types flottants :

- `float32` : 32 bits (~6-9 chiffres de précision).
- `float64` : 64 bits (~15-17 chiffres de précision). C'est le type flottant par défaut.

Comme on peut le voir, les deux types peuvent stocker un nombre différent de chiffres. Cela signifie que l'essai de stockage de PI dans un `float32` ne stockera que les 6 à 9 premiers chiffres (le dernier chiffre étant arrondi).

Par défaut, Go utilisera `float64` pour les nombres flottants, sauf si le nombre flottant est :

1. assigné à une variable de type `float32`, ou
2. retourné par une fonction avec le type de retour `float32`, ou
3. passé comme argument à la fonction `float32()`.

Le paquet `math` contient de nombreuses fonctions mathématiques utiles.

## Instructions

Dans cet exercice, vous travaillerez avec des comptes d'épargne.
Chaque année, le solde de votre compte d'épargne est mis à jour en fonction de son taux d'intérêt.
Le taux d'intérêt que votre banque vous donne dépend du montant d'argent dans votre compte (son solde) :

- 3.213% pour un solde inférieur à `0` dollars (le solde devient plus négatif).
- 0.5% pour un solde supérieur ou égal à `0` dollars et inférieur à `1000` dollars.
- 1.621% pour un solde supérieur ou égal à `1000` dollars et inférieur à `5000` dollars.
- 2.475% pour un solde supérieur ou égal à `5000` dollars.

Vous avez quatre tâches, chacune traitant votre solde et son taux d'intérêt.

## 1. Calculer le taux d'intérêt

Implémentez la fonction `InterestRate()` pour calculer le taux d'intérêt en fonction du solde spécifié :

```go
InterestRate(200.75)
// => 0.5
```

Notez que la valeur retournée est un `float32`.

## 2. Calculer l'intérêt

Implémentez la fonction `Interest()` pour calculer l'intérêt en fonction du solde spécifié :

```go
Interest(200.75)
// => 1.003750
```

Notez que la valeur retournée est un `float64`.

## 3. Calculer la mise à jour annuelle du solde

Implémentez la fonction `AnnualBalanceUpdate()` pour calculer la mise à jour annuelle du solde, en tenant compte du taux d'intérêt :

```go
AnnualBalanceUpdate(200.75)
// => 201.75375
```

Notez que la valeur retournée est un `float64`.

## 4. Calculer les années avant d'atteindre le solde souhaité

Implémentez la fonction `YearsBeforeDesiredBalance()` pour calculer le nombre minimum d'années nécessaires pour atteindre le solde souhaité, en tenant compte du fait que chaque année, les intérêts sont ajoutés au solde.
Cela signifie que le solde après un an est : solde initial + intérêt pour solde initial.
Le solde après la deuxième année est : solde après un an + intérêt pour solde après un an.
Et ainsi de suite, jusqu'à ce que le solde de l'année en cours soit supérieur ou égal au solde cible.

```go
balance := 200.75
targetBalance := 214.88
YearsBeforeDesiredBalance(balance, targetBalance)
// => 14
```

Notez que la valeur retournée est un `int`.

## Source

### Créé par

- @erikschierboom