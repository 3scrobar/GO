# Réservation pour la Beauté

Bienvenue dans « Réservation pour la Beauté » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

Un [`Time`][time] dans Go est un type décrivant un moment dans le temps. Les informations de date et d'heure peuvent être consultées, comparées et manipulées via ses méthodes, mais il existe également certaines fonctions appelées sur le package `time` lui-même. La date et l'heure actuelles peuvent être récupérées via la fonction [`time.Now`][now].

La fonction [`time.Parse`][parse] analyse les chaînes en valeurs de type `Time`. Go a une manière particulière de définir la mise en page que vous attendez pour l'analyse. Vous devez écrire un exemple de mise en page en utilisant les valeurs de cet horodatage spécial :
`Mon Jan 2 15:04:05 -0700 MST 2006`.

Par exemple:

```go
import "time"

func parseTime() time.Time {
    date := "Tue, 09/22/1995, 13:00"
    layout := "Mon, 01/02/2006, 15:04"

    t, err := time.Parse(layout,date) // time.Time, error
}

// => 1995-09-22 13:00:00 +0000 UTC
```

La méthode [`Time.Format()`][format] renvoie une représentation sous forme de chaîne de l'heure. Tout comme avec la fonction `Parse`, la disposition cible est à nouveau définie via un exemple qui utilise les valeurs de l'horodatage spécial.

Par exemple:

```go
import (
    "fmt"
    "time"
)

func main() {
    t := time.Date(1995,time.September,22,13,0,0,0,time.UTC)
    formattedTime := t.Format("Mon, 01/02/2006, 15:04") // string
    fmt.Println(formattedTime)
}

// => Fri, 09/22/1995, 13:00
```

## Options de mise en page

Pour une mise en page personnalisée, utilisez une combinaison de ces options. Dans Go, la date et l'horodatage prédéfinis [constantes de format] [const] sont également disponibles.

| Temps | Options |
| ----------- | ---------------------------------------------- |
| Année | 2006 ; 06 |
| Mois | Jan ; Janvier ; 01 ; 1 |
| Jour | 02 ; 2 ; \_2 (Pour le 0 précédent) |
| Jour de la semaine | Lun ; Lundi |
| Heure | 15 (format heure 24 heures) ; 3 ; 03 (AM ou PM) |
| Procès-verbal | 04 ; 4 |
| Deuxième | 05 ; 5 |
| Marque AM/PM | PM |
| Jour de l'année | 002 ; \__\_2 |

Le type `time.Time` dispose de différentes méthodes pour accéder à une heure particulière. par ex. Heure : [`Time.Hour()`][heure] , Mois : [`Time.Month()`][mois]. Pour en savoir plus sur la façon dont cela fonctionne, consultez la [documentation officielle][time].

Le [`time`][time] inclut un autre type, [`Duration`][duration], représentant le temps écoulé, ainsi que la prise en charge des emplacements/fuseaux horaires, des minuteries et d'autres fonctionnalités associées qui seront couvertes dans un autre concept.

[time]: https://golang.org/pkg/time/#Time
[now]: https://golang.org/pkg/time/#Now
[const]: https://pkg.go.dev/time#pkg-constants
[format]: https://pkg.go.dev/time#Time.Format
[hour]: https://pkg.go.dev/time#Time.Hour
[month]: https://pkg.go.dev/time/#Time.Month
[duration]: https://pkg.go.dev/time#Duration
[parse]: https://golang.org/pkg/time/#Parse
[article]: https://www.pauladamsmith.com/blog/2011/05/go_time.html

## Instructions

Dans cet exercice, vous travaillerez sur un planificateur de rendez-vous pour un salon de beauté ouvert le 15 septembre 2012.

Vous avez cinq tâches, qui impliqueront toutes des dates de rendez-vous.

## 1. Analyser la date du rendez-vous

Implémentez la fonction `Schedule` pour analyser une représentation textuelle d'une date de rendez-vous dans le format `time.Time` correspondant :

```go
Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC
```

## 2. Vérifiez si un rendez-vous est déjà passé

Implémentez la fonction `HasPassed` qui prend une date de rendez-vous et vérifie si le rendez-vous a eu lieu quelque part dans le passé :

```go
HasPassed("July 25, 2019 13:45:00")
// => true
```

## 3. Vérifiez si le rendez-vous est dans l'après-midi

Implémentez la fonction `IsAfternoonAppointment` qui prend une date de rendez-vous et vérifie si le rendez-vous est dans l'après-midi (>= 12h00 et < 18h00) :

```go
IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// => true
```

## 4. Décrivez l'heure et la date du rendez-vous

Implémentez la fonction `Description` qui prend une date de rendez-vous et renvoie une description de cette date et heure :

```go
Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
```

## 5. Retourner la date anniversaire de l'ouverture du salon

Implémentez la fonction `AnniversaryDate` qui renvoie la date anniversaire d'ouverture du salon pour l'année en cours en UTC.

En supposant que l’année en cours soit 2020 :

```go
AnniversaryDate()

// => 2020-09-15 00:00:00 +0000 UTC
```

**Remarque :** la valeur de retour est un `time.Time` et l'heure de la journée n'a pas d'importance.

## Source

### Créé par

- @jamessouth
