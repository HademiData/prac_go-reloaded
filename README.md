## Practice for my 01-EDU-GORELOADED PROJECT

**go-reloaded** is a Go-based text processing program that reads a text file, applies a set of transformation rules, and writes the corrected result to another file.

The program modifies text by applying formatting rules such as punctuation correction, casing adjustments, number conversions, article corrections, and bracket handling.

This project was built to practice and improve skills in **Go programming**, **string manipulation**, **file handling**, and **modular code organization**.

---


## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Files and Their Functions](#files-and-their-functions)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Testing](#testing)
- [Author](#author)

---

## Features ✨

- Automatic punctuation formatting
- Case transformations (**upper**, **lower**, **capitalize**)
- Number formatting
- Article correction (**a → an** when appropriate)
- Bracket handling
- Modular Go code structure
- File input and output support

---


## Project Structure 📁


```
go-reloaded/
│
├── main.go
├── go.mod
├── utils/
│   ├── app/
│   │   └── App.go
│   │
│   ├── helperFunctions/
│   │   ├── capitalize.go
│   │   ├── groupOfPunctuations.go
│   │   ├── lower.go
│   │   └── upper.go
│   │
│   ├── adjustPunctuations.go
│   ├── adjustQuotations.go
│   ├── article.go
│   ├── formatCasing.go
│   └── replaceBInAndHex.go
│
├── results.txt
├── sample.txt
├── README.md
│
└── tests/
    ├── adjustPunctuations_test.go
    ├── adjustQuotations_test.go
    ├── capitalize_test.go
    ├── formatCasing_test.go
    ├── lowerCasing_test.go
    ├── modifyArticle_test.go
    ├── replaceBinAndHex_test.go
    └── upperCase_test.go

```

## Files and Their Functions
**main.go** – Entry point of the program  
**App.go** – Handles reading from and writing to files  
**FormatCasing.go** – Handles case transformations  
**replaceBinAndHex_test.go** – Handles binary and hexadecimal number conversion  
**adjustPunctuation.go and adjustQuotations.go** – Handles punctuation corrections  
**articles.go** – Handles article corrections (`a` / `an`)  
**tests/** – Contains unit tests for the project  

---

## Installation ⚙️

1. Make sure **Go** is installed.

Check with:

```bash
go version
```

2. Clone the repository:

```bash
git clone https://github.com/your-username/go-reloaded.git
cd go-reloaded
```

## Usage ▶️
Run the program with:

```bash
go run . sample.txt result.txt
```
Where:

sample.txt → input file

result.txt → output file

The program processes the text and writes the corrected result to the output file.


## Example

```bash

$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
$ cat sample.txt
There is no greater agony than bearing a untold story inside you.
$ go run . sample.txt result.txt
$ cat result.txt
There is no greater agony than bearing an untold story inside you.
$ cat sample.txt
Punctuation tests are ... kinda boring ,what do you think ?
$ go run . sample.txt result.txt
$ cat result.txt
Punctuation tests are... kinda boring, what do you think?
```

## Testing 🧪

Run all tests with:

```bash
go test ./tests
```
This will execute all test files located in the tests directory.

## Author 👨‍💻

Afolabi Adewale