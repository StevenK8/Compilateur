# Compilateur

Steven Kerautret & Silvain Théréné


## Compilation :

go build

[![Build Status](https://travis-ci.org/StevenK8/Compilateur.png?branch=master)](https://travis-ci.org/StevenK8/Compilateur)

## Flags:

-file   :   Path du fichier en entrée (-file='test.txt')

-o      :   Path du fichier de sortie (-o='test.out')

-h      :   Affiche l'aide


## Exécution:

```sh
./Compilateur.exe -file='test.txt' -o='test.out'    (Windows)
```

```sh
./Compilateur -file='test.txt' -o='test.out'    (Linux)
```

## TO DO (For 28 / 11 )
### Tester 
* [ ] Exécute (compile (Texte)) = Out
* [ ] Demander précieuse, tout voir le reste du cours.https://cdn.discordapp.com/attachments/711219342985134090/779660615543029790/unknown.png
* [ ] Sortie brut = Code MSM
* [X] Diviser en sortie (Test_mult_1.out, Test_mult_2.out, ….)

### Objectif
* [x] Analyseur lexical
* [x] Atomes (Maths)
* [x] Expressions (sauf affectation)
* [x] Variables (déclaration / affectation)
* [x] Table des symboles
* [x] Analyse sémantique
* [x] Conditionelles
* [-] Boucles (break)
* [-] Fonctions (définitions)
* [-] Fonctions (Appels)

* [-] Pointeur
* [-] Tableaux
* [-] Runtime ) pour print
