# Analyse de fichiers journaux

Bienvenue dans le défi Analyse de fichiers journaux sur la piste Go d'Exercism.
Si vous avez besoin d'aide pour exécuter les tests ou soumettre votre code, consultez `HELP.md`.
Si vous êtes bloqué sur l'exercice, consultez `HINTS.md`, mais essayez d'abord de le résoudre sans l'utiliser :)

## Introduction

Le paquet [regexp][package-regexp] offre un support pour les expressions régulières en Go.

## Syntaxe 

La [syntaxe][regexp-syntax] des expressions régulières acceptées est la même syntaxe générale utilisée par Perl, Python et d'autres langages. 

Les modèles de recherche et les textes d'entrée sont interprétés comme UTF-8.  

Lors de l'utilisation de backticks (\`) pour créer des chaînes, les antislashs (`\`) n'ont aucune signification particulière et ne marquent pas le début de caractères spéciaux comme les tabulations `\t` ou les sauts de ligne `\n` :

```go
"\t\n" // chaîne littérale régulière avec 2 caractères : une tabulation et un saut de ligne
`\t\n`// chaîne littérale brute avec 4 caractères : deux antislashs, un 't' et un 'n'
```

Pour cette raison, l'utilisation de backticks est souhaitable pour les expressions régulières,
car cela signifie que nous n'avons pas besoin d'échapper les antislashs :

```go
"\\" // chaîne avec un seul antislash
`\\` // chaîne avec 2 antislashs
```

## Compilation des modèles - type `RegExp`

Pour utiliser une expression régulière, nous devons d'abord compiler le modèle de chaîne.
La compilation signifie ici prendre le modèle de chaîne de l'expression régulière et le convertir en une représentation interne plus facile à utiliser.
Nous n'avons besoin de compiler chaque modèle qu'une seule fois, après quoi nous pouvons utiliser la version compilée de l'expression régulière de nombreuses fois.
Le type `regexp.Regexp` représente une expression régulière compilée.
Nous pouvons compiler un modèle de chaîne en `regexp.Regexp` en utilisant la fonction `regexp.Compile`.
Cette fonction retourne `nil` et une erreur en cas d'échec de la compilation :

```go
re, err := regexp.Compile(`(a|b)+`)
fmt.Println(re, err) // => (a|b)+ <nil>
re, err = regexp.Compile(`a|b)+`)
fmt.Println(re, err) // => <nil> erreur lors de l'analyse de regexp : inattendu ) : `a|b)+`
```

La fonction `MustCompile` est une alternative pratique à `Compile` : 

```go 
re = regexp.MustCompile(`[a-z]+\d*`)
```

Avec cette fonction, il n'est pas nécessaire de gérer une erreur. 

~~~~exercism/caution
 `MustCompile` ne doit être utilisée que lorsque nous sommes sûrs que le modèle se compile, sinon le programme panique.
 ~~~~
 
 ## Méthodes d'expression régulière
 
Il existe 16 méthodes de `Regexp` qui correspondent à une expression régulière et identifient le texte mis en correspondance.
Leurs noms correspondent à cette expression régulière :

```text
Find(All)?(String)?(Submatch)?(Index)?
```

* Si `All` est présent, la routine fait correspondre les correspondances successives non chevauchantes de l'expression entière.
* Si `String` est présent, l'argument est une chaîne ; sinon c'est une tranche d'octets ; les valeurs de retour sont ajustées en conséquence. 
* Si `Submatch` est présent, la valeur de retour est une tranche identifiant les sous-correspondances successives de l'expression.
* Si `Index` est présent, les correspondances et les sous-correspondances sont identifiées par des paires d'index d'octets dans la chaîne d'entrée.

Il existe également des méthodes pour :

* remplacer les correspondances d'expressions régulières par des chaînes de remplacement et
* diviser les chaînes séparées par des expressions régulières.

En tout, le paquet `regexp` définit plus de 40 fonctions et méthodes.
Nous démontrerons l'utilisation de quelques méthodes ci-dessous.
Veuillez consulter la [documentation API][package-regexp] pour plus de détails sur ces fonctions et d'autres.
 
### Exemples de `MatchString` 

La méthode `MatchString` indique si une chaîne contient une correspondance d'une expression régulière.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
b = re.MatchString("[a12]")       // => true
b = re.MatchString("12abc34(ef)") // => true
b = re.MatchString(" abc!")       // => true
b = re.MatchString("123 456")     // => false    
```

### Exemples de `FindString` 

La méthode `FindString` retourne une chaîne contenant le texte de la correspondance la plus à gauche de l'expression régulière.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
s = re.FindString("[a12]")       // => "a12"
s = re.FindString("12abc34(ef)") // => "abc34"
s = re.FindString(" abc!")       // => "abc"
s = re.FindString("123 456")     // => ""
```

### Exemples de `FindStringSubmatch`

La méthode `FindStringSubmatch` retourne une tranche de chaînes contenant le texte de la correspondance la plus à gauche de l'expression régulière et les correspondances, le cas échéant, de ses sous-expressions.
Cela peut être utilisé pour identifier les chaînes correspondant aux groupes de capture.
Une valeur de retour de `nil` indique l'absence de correspondance.

```go 
re = regexp.MustCompile(`[a-z]+(\d*)`)
sl = re.FindStringSubmatch("[a12]")       // => []string{"a12","12"}
sl = re.FindStringSubmatch("12abc34(ef)") // => []string{"abc34","34"}
sl = re.FindStringSubmatch(" abc!")       // => []string{"abc",""}
sl = re.FindStringSubmatch("123 456")     // => <nil>
```

### Exemples de `ReplaceAllString`

La méthode `re.ReplaceAllString(src,repl)` retourne une copie de `src`, en remplaçant les correspondances de l'expression régulière `re` par la chaîne de remplacement `repl`.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
s = re.ReplaceAllString("[a12]", "X")       // => "[X]"
s = re.ReplaceAllString("12abc34(ef)", "X") // => "12X(X)"
s = re.ReplaceAllString(" abc!", "X")       // => " X!"
s = re.ReplaceAllString("123 456", "X")     // => "123 456"
```
 
 ### Exemples de `Split`
 
La méthode `re.Split(s,n)` divise un texte `s` en sous-chaînes séparées par l'expression et retourne une tranche des sous-chaînes entre ces correspondances d'expression.
Le compte `n` détermine le nombre maximal de sous-chaînes à retourner.
Si `n<0`, la méthode retourne toutes les sous-chaînes.

```go
re = regexp.MustCompile(`[a-z]+\d*`)
sl = re.Split("[a12]", -1)      // => []string{"[","]"}
sl = re.Split("12abc34(ef)", 2) // => []string{"12","(ef)"}
sl = re.Split(" abc!", -1)      // => []string{" ","!"}
sl = re.Split("123 456", -1)    // => []string{"123 456"}
```
  
[package-regexp]: https://pkg.go.dev/regexp
[regexp-syntax]: https://pkg.go.dev/regexp/syntax

## Instructions

Cet exercice traite de l'analyse de fichiers journaux.

Suite à un examen de sécurité récent, vous avez été invité à nettoyer les fichiers journaux archivés de l'organisation.

Toutes les chaînes transmises aux fonctions sont garanties d'être non-nulles et sans espaces de début ou de fin.

## 1. Identifier les lignes de journal corrompues

Vous avez besoin d'une idée du nombre de lignes de journal dans votre archive qui ne sont pas conformes aux normes actuelles.
Vous pensez qu'un test simple révèle si une ligne de journal est valide.
Pour être considérée comme valide, une ligne doit commencer par l'une des chaînes suivantes :

- [TRC]
- [DBG]
- [INF]
- [WRN]
- [ERR]
- [FTL]

Implémentez la fonction `IsValidLine` pour retourner `false` si une chaîne n'est pas valide sinon `true`.

```go 
IsValidLine("[ERR] A good error here")
// => true
IsValidLine("Any old [ERR] text")
// => false
IsValidLine("[BOB] Any old text")
// => false
```

## 2. Diviser la ligne de journal

Une nouvelle équipe a rejoint l'organisation, et vous découvrez que ses fichiers journaux utilisent un séparateur étrange pour les "champs".
Au lieu de quelque chose de sensé comme deux points ":" ils utilisent une chaîne telle que "<--->" ou "<=>" (parce que c'est plus joli) en fait n'importe quelle chaîne qui a un premier caractère "<" et un dernier caractère ">" et n'importe quelle combinaison des caractères suivants "~", "\*", "=" et "-" entre les deux.

Implémentez la fonction `SplitLogLine` qui prend une ligne et retourne un tableau de chaînes dont chacune contient un champ.

```go
SplitLogLine("section 1<*>section 2<~~~>section 3")
// => []string{"section 1", "section 2", "section 3"},
```

## 3. Compter le nombre de lignes contenant `password` dans du texte entre guillemets

L'équipe a besoin de connaître les références aux mots de passe dans le texte entre guillemets afin qu'elles puissent être examinées manuellement.

Implémentez la fonction `CountQuotedPasswords` pour donner une indication de l'ampleur probable de l'exercice manuel.

Identifiez les lignes de journal où la chaîne "password", qui peut être en n'importe quelle combinaison de majuscules ou minuscules, est entourée de guillemets.
Vous devez tenir compte de la possibilité d'un contenu supplémentaire entre les guillemets avant et après "password".
Chaque ligne contient au maximum deux guillemets.

Les lignes transmises à la routine peuvent être ou non valides selon la définition de la tâche 1.
Nous les traitons de la même manière, qu'elles soient valides ou non.

```go
lines := []string{
    `[INF] passWord`, // contient 'password' mais pas entouré de guillemets
    `"passWord"`,  // comptez celle-ci
    `[INF] User saw error message "Unexpected Error" on page load.`, // ne contient pas 'password'
    `[INF] The message "Please reset your password" was ignored by the user`, // comptez celle-ci 
}
// => 2
```

## 4. Supprimer les artefacts du journal

Vous avez découvert que certains traitements en amont des journaux ont dispersé le texte "end-of-line" suivi d'un numéro de ligne (sans espace intermédiaire) partout dans les journaux.

Implémentez la fonction `RemoveEndOfLineText` pour prendre une chaîne et supprimer le texte de fin de ligne et retourner une chaîne "propre".

Les lignes ne contenant pas de texte de fin de ligne doivent être retournées inchangées.

Supprimez simplement la chaîne de fin de ligne. N'essayez pas d'ajuster les espaces.

```go
RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27")
// => "[INF]  Network Failure "
```

## 5. Étiqueter les lignes avec les noms d'utilisateurs

Vous avez remarqué que certaines lignes de journal incluent des phrases qui font référence aux utilisateurs.
Ces phrases contiennent toujours la chaîne `"User"`, suivie d'un ou plusieurs espaces, puis d'un nom d'utilisateur.
Vous décidez d'étiqueter de telles lignes.

Implémentez une fonction `TagWithUserName` qui traite les lignes de journal :

- Les lignes qui ne contiennent pas la chaîne `"User "` restent inchangées.
- Pour les lignes qui contiennent la chaîne `"User "`, préfixez la ligne avec `[USR]` suivi du nom d'utilisateur.
 
Par exemple :

```go
result := TagWithUserName([]string{
    "[WRN] User James123 has exceeded storage space.",
    "[WRN] Host down. User   Michelle4 lost connection.",
    "[INF] Users can login again after 23:00.",
    "[DBG] We need to check that user names are at least 6 chars long.",
     
}) 
// => []string {
//  "[USR] James123 [WRN] User James123 has exceeded storage space.",
//  "[USR] Michelle4 [WRN] Host down. User   Michelle4 lost connection.",
//  "[INF] Users can login again after 23:00.",
//  "[DBG] We need to check that user names are at least 6 chars long."
// }
```

Vous pouvez supposer que : 

- Les noms d'utilisateurs sont suivis d'au moins un caractère d'espace dans le journal.
- Il y a au maximum une occurrence de la chaîne `"User "` dans chaque ligne.
- Les noms d'utilisateurs sont des chaînes non vides qui ne contiennent pas d'espace.

## Source

### Créé par

- @norbs57

### Contribué par

- @eklatzer