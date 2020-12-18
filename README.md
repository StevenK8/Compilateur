# Compilateur

<div>
  <p>
    <img src="https://img.shields.io/badge/Author-Steven Kerautret-yellow.svg" />
    <img src="https://img.shields.io/badge/-Silvain Théréné-yellow.svg" />
  </p>
  <p>
    <img src="https://img.shields.io/badge/Status-on development-success.svg" />
  </p>
  <p>
    <a href="https://github.com/StevenK8/Compilateur/actions">
      <img src="https://img.shields.io/badge/Build-Failed-critical.svg" />
    </a>
    <img src="https://img.shields.io/badge/Code coverage-0 %25-critical.svg" />
  </p>
  
  <p>
    <img src="https://img.shields.io/badge/Total tests-6-important.svg" />
    <img src="https://img.shields.io/badge/Tests passed-0-critical.svg" />
    <img src="https://img.shields.io/badge/Test quality-0 %25-critical.svg" />
  </p>
  <p>
    <img src="https://img.shields.io/badge/langage-golang-7fd5ea.svg" />
    <img src="https://img.shields.io/badge/Go-1.15.5-informational.svg" />
    <img src="https://img.shields.io/badge/platform-linux-lightgray.svg" />
    <img src="https://img.shields.io/badge/-windows-lightgray.svg" />
  </p>
  <p>
    <img src="https://img.shields.io/badge/IDE Used-Visual Studio Code-informational.svg" />
    <img src="https://img.shields.io/badge/-Goland Jetbrains-informational.svg" />
  </p>
</div>

## Compilation :

```sh
go build
```

## Flags:

```sh
-file   :   Path du fichier en entrée (-file='test.txt')

-o      :   Path du fichier de sortie (-o='test.out')

-h      :   Affiche l'aide
```


## Exécution:

```sh
./Compilateur.exe -file='test.txt' -o='test.out'    (Windows)
```

```sh
./Compilateur -file='test.txt' -o='test.out'    (Linux)
```


### Features
* [x] Analyseur lexical
* [x] Atomes (Maths)
* [x] Expressions (sauf affectation)
* [x] Variables (déclaration / affectation)
* [x] Table des symboles
* [x] Analyse sémantique
* [x] Conditionelles
* [x] Boucles (break)
* [x] Fonctions (définitions)
* [x] Fonctions (Appels)
* [x] Pointeur
* [x] Tableaux
* [x] Runtime pour print
