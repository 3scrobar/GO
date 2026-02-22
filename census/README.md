# Recensement

Bienvenue dans « Recensement » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Go n'a pas de concept de vide, nul ou indéfini pour les valeurs de variable. Les variables déclarées sans valeur initiale explicite ont par défaut la valeur zéro pour leur type respectif.

La valeur zéro pour les types primitifs tels que les booléens, les types numériques et les chaînes est respectivement `false`, `0` et `""`.

L'identifiant `nil`, signifiant zéro, est la valeur zéro pour les types plus complexes tels que les pointeurs, les fonctions, les interfaces, les tranches, les canaux et les cartes.

Le tableau suivant détaille la valeur zéro pour les types Go.

| Tapez | Valeur zéro |
| --------- | ---------- |
| booléen | `false` |
| numérique | `0` |
| chaîne | `""` |
| pointeur | `nil` |
| fonction | `nil` |
| interfaces | `nil` |
| tranche | `nil` |
| canal | `nil` |
| carte | `nil` |

Vous avez peut-être remarqué que les types de structures sont absents du tableau ci-dessus. En effet, la valeur zéro d'un type struct dépend de ses champs. Les structures sont définies sur leur valeur zéro lorsque tous leurs champs sont définis sur leur valeur zéro respective.

## Instructions

Il vous incombe de préparer le système informatique de la ville pour un prochain recensement. Plus précisément, vous êtes responsable du programme qui traitera les données des agents recenseurs.

Le programme doit être capable de créer un nouveau résident dans le système lorsqu'il reçoit les informations d'un résident. De plus, vous créerez des fonctions qui garantiront que les informations requises sont présentes dans les données du résident et supprimerez les données d'un résident. Enfin, vous compterez les résidents pour fournir un décompte précis.

## 1. Créer un nouveau résident

Lorsqu'un agent du recensement collecte les informations d'un résident, il doit enregistrer ce résident en saisissant son nom, son âge et son adresse dans le système.

Implémentez la fonction `NewResident` qui accepte trois arguments :

- Le nom du résident.
- L'âge du résident.
- L'adresse du résident.

La fonction doit renvoyer un pointeur vers une structure `Resident` qui contient ces informations.

```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St."}

NewResident(name, age, address)
// => &{Matthew Sanabria 29 map[street:Main St.]}
```

## 2. Validez les informations requises

Les résidents peuvent être réticents à fournir des données personnelles aux agents recenseurs. Dans ces cas, il est nécessaire de déterminer si le résident a fourni les informations requises pour être pris en compte lors du recensement.

Pour être compté, un résident doit fournir une valeur non nulle pour son nom et une adresse qui contient une valeur non nulle pour la clé `street`. Toutes les autres informations, comme l'âge du résident, sont facultatives. Implémentez la méthode `HasRequiredInfo` qui renvoie un booléen indiquant si le résident a fourni les informations requises.

```go
name := "Matthew Sanabria"
age := 0
address := make(map[string]string)

resident := NewResident(name, age, address)

resident.HasRequiredInfo()
// => false
```

## 3. Supprimer les informations du résident

La vie avance vite et des erreurs se produisent. Un résident peut quitter la ville. Un recenseur peut commettre des erreurs lors de la collecte de données. Dans ces cas-là, il est nécessaire d'avoir la possibilité de supprimer les données d'un résident pour qu'il ne soit pas comptabilisé.

Implémentez la méthode `Delete` qui définit tous les champs du résident à leur valeur zéro.

```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St."}

resident := NewResident(name, age, address)

fmt.Println(resident)
// Output: &{Matthew Sanabria 29 map[street:Main St.]}

resident.Delete()

fmt.Println(resident)
// Output: &{ 0 map[]}
```

## 4. Comptez les résidents

Maintenant que le système prend en charge les données de recensement, il est temps d'effectuer le recensement et de compter les habitants !

Implémentez la fonction `Count` qui accepte un argument :

- Une tranche de pointeurs vers les structures `Resident`.

La fonction doit renvoyer un entier indiquant le nombre de résidents recensés lors du recensement. Un résident ne peut être compté que s’il a fourni les informations requises au recenseur.

```go
name1 := "Matthew Sanabria"
age1 := 29
address1 := map[string]string{"street": "Main St."}

resident1 := NewResident(name1, age1, address1)

name2 := "Rob Pike"
age2 := 0
address2 := make(map[string]string)

resident2 := NewResident(name2, age2, address2)

residents := []*Resident{resident1, resident2}

Count(residents)
// => 1
```

## Source

### Créé par

- @jamessouth
- @sudomateo

### Contribué à par

- @tehsphinx
- @eklatzer
