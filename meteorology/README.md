# Météorologie

Bienvenue dans « Météorologie » sur le parcours Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans aide :).

## Introduction

[Stringer][stringer-interface] est une interface permettant de définir le format de chaîne des valeurs.

L'interface se compose d'une seule méthode `String` :
 
```go
type Stringer interface {
    String() string
}
```
 
Les types qui souhaitent implémenter cette interface doivent avoir une méthode `String()` qui renvoie une représentation sous forme de chaîne conviviale du type. Le package [fmt][fmt-package] (et bien d'autres) recherchera cette méthode pour formater et imprimer les valeurs.

## Exemple : Distances

Supposons que nous travaillions sur une application qui traite les distances géographiques mesurées dans différentes unités.
Nous avons défini les types `DistanceUnit` et `Distance` comme suit :
 
```go 
type DistanceUnit int

const (
	Kilometer    DistanceUnit = 0
	Mile         DistanceUnit = 1
)
 
type Distance struct {
	number float64
	unit   DistanceUnit
} 
```

Dans l'exemple ci-dessus, `Kilometer` et `Mile` sont des constantes de type `DistanceUnit`.

Ces types n'implémentent pas l'interface `Stringer` car ils ne disposent pas de la méthode `String`.
Par conséquent, les fonctions `fmt` imprimeront les valeurs `Distance` en utilisant le "format par défaut" de Go :

```go
mileUnit := Mile
fmt.Sprint(mileUnit)
// => 1
// The result is '1' because that is the underlying value of the 'Mile' constant (see constant declarations above) 

dist := Distance{number: 790.7, unit: Kilometer}
fmt.Sprint(dist)
// => {790.7 0}
// not a very useful output!
```

Afin de rendre la sortie plus informative, nous implémentons l'interface `Stringer` pour les types `DistanceUnit` et `Distance` en ajoutant une méthode `String` à chaque type :

```go
func (sc DistanceUnit) String() string {
	units := []string{"km", "mi"}
	return units[sc]
}

func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
}
```
 
Les fonctions du package `fmt` appelleront ces méthodes lors du formatage des valeurs `Distance` :

```go
kmUnit := Kilometer
kmUnit.String()
// => km

mileUnit := Mile
mileUnit.String()
// => mi

dist := Distance{
	number: 790.7,
	unit: Kilometer,
}
dist.String()
// => 790.7 km
```

[stringer-interface]: https://pkg.go.dev/fmt#Stringer
[fmt-package]: https://pkg.go.dev/fmt

## Instructions

Votre équipe travaille sur une application météorologique.
Ils ont défini une API avec différents types et constantes représentant les données météorologiques, voir fichier `meteorology.go`.
  
Votre tâche consiste à ajouter des méthodes `String` appropriées à tous les types afin qu'ils implémentent l'interface `Stringer`.

## 1. Implémentez l'interface `Stringer` pour le type `TemperatureUnit`

Après quelques discussions, l'équipe a convenu que l'unité de température serait soit `Celsius`, soit `Fahrenheit`. Les valeurs doivent être formatées comme indiqué dans les exemples ci-dessous.

Faites en sorte que le type `TemperatureUnit` implémente l'interface `Stringer` en y ajoutant une méthode `String`. Cette méthode doit renvoyer la chaîne `"°C"` si l'unité de température est Celsius ou `"°F"` si l'unité de température est Fahrenheit.

```go
celsiusUnit := Celsius
fahrenheitUnit := Fahrenheit

celsiusUnit.String()
// => °C
fahrenheitUnit.String()
// => °F
fmt.Sprint(celsiusUnit)
// Output: °C
```

## 2. Implémentez l'interface `Stringer` pour le type `Temperature`

Les valeurs de température sont constituées d'un nombre entier et d'une unité de température. Ils doivent être formatés comme dans les exemples ci-dessous.

Pour que cela se produise, faites en sorte que le type `Temperature` implémente l'interface `Stringer` en y ajoutant une méthode `String`. Cette méthode doit renvoyer une chaîne avec la valeur numérique de la température et l'unité de température séparées par un espace : `<temperature> <unit>` :


```go
celsiusTemp := Temperature{
    degree: 21,
    unit: Celsius,
}
celsiusTemp.String()
// => 21 °C
fmt.Sprint(celsiusTemp)
// Output: 21 °C

fahrenheitTemp := Temperature{
    degree: 75,
    unit: Fahrenheit,
}
fahrenheitTemp.String()
// => 75 °F
fmt.Sprint(fahrenheitTemp) 
// Output: 75 °F
```

## 3. Implémentez l'interface `Stringer` pour le type `SpeedUnit`

Après de longues discussions, l'équipe a convenu que l'unité de vitesse du vent serait soit `KmPerHour`, soit `MilesPerHour`. Les valeurs doivent être formatées comme les exemples ci-dessous.

Pour que cela se produise, faites en sorte que le type `SpeedUnit` implémente l'interface `Stringer` en y ajoutant une méthode `String`. Cette méthode doit renvoyer la chaîne `"km/h"` si l'unité de vitesse est en kilomètres par heure ou `"mph"` si l'unité de vitesse est en miles par heure :


```go 
mphUnit := MilesPerHour
mphUnit.String()
// => mph
fmt.Sprint(mphUnit)
// Output: mph

kmhUnit := KmPerHour
kmhUnit.String()
// => km/h
fmt.Sprint(kmhUnit)
// Output: km/h
```

## 4. Implémentez l'interface `Stringer` pour `Speed`

Les valeurs de vitesse du vent se composent d'un nombre entier et d'une unité de vitesse. Ils doivent être formatés comme dans l’exemple ci-dessous.

Pour que cela se produise, faites en sorte que le type `Speed` implémente l'interface `Stringer` en y ajoutant une méthode `String`. Cette méthode doit renvoyer une chaîne avec la valeur numérique de la vitesse et l'unité de vitesse séparées par un espace : `<speed> <unit>` :

```go 
windSpeedNow := Speed{
    magnitude: 18,
    unit: KmPerHour,
}
windSpeedNow.String(windSpeedNow)
// => 18 km/h
fmt.Sprintf(windSpeedNow)
// Output: 18 km/h

windSpeedYesterday := Speed{
    magnitude: 22,
    unit: MilesPerHour,
}
windSpeedYesterday.String(windSpeedYesterday)
// => 22 mph
fmt.Sprint(windSpeedYesterday)
// Output: 22 mph
```

## 5. Implémentez l'interface `Stringer` pour le type `MeteorologyData`

Les données météorologiques précisent l'emplacement, la température, la direction du vent et la vitesse du vent.
et l'humidité. Il doit être formaté comme dans l'exemple ci-dessous :

Pour que cela se produise, faites en sorte que le type `MeteorologyData` implémente l'interface `Stringer` en y ajoutant une méthode `String`. Cette méthode doit renvoyer les données météorologiques au format suivant :

```
<location>: <temperature>, Wind <wind_direction> at <wind_speed>, <humidity>% Humidity
```

```go 
sfData := MeteorologyData{
    location: "San Francisco",
    temperature: Temperature{
        degree: 57,
        unit: Fahrenheit
    },
    windDirection: "NW",
    windSpeed: Speed{
        magnitude: 19,
        unit: MilesPerHour
    },
    humidity: 60
}

sfData.String()
// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
fmt.Sprint(sfData) 
// Output: San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
```

## Source

### Créé par

- @norbs57

### Contribué à par

- @andrerfcsantos
- @eklatzer
