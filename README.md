# A Simple Start

5 Important Questions:

- How do we run the code in our project?

  - command terminal: go run main.go

  - GO CLI:

    - go build (compiles a bunch of go source code files)
    - go run (compiles and executes one or two files)
    - go fmt (formats all the code in each file in the current directory)
    - go install (compiles and "installs" a package)
    - go get (downloads the raw source ode of someone else's package)
    - go test (runs any tests associated with the current project)

- What does package main mean?

  - Package == Project == Workspace (a collection of common source code files)

    - A package can have many related files inside it (each file ending with extension GO)
      Only requirement for every file inside of a package is that the very first line of each file must declare the package it belongs to

    - Types of Packages:

      - Executable (Generates a file that we can run)

        - package main (defines a package that can be compiled then "executed")
          Note: Must have a func called main

      - Reusable (Code used as "helpers" e.g. dependencies. Good place to put reusable logic)

        - package calculator/uploader (defines a package that can be used as a dependency)

    - The name of the package - executable/dependency type (E.g. package main - main is executable type)

    - package main ---> go build ---> main.exe
      <br/>
      (if we ran this file, function named "main" would be automatically run)

    - Package name (main) is sacred as we use only when making a package for runnable file

- What does import "fmt" mean?

  - import fmt means give my package main access to all of the code and all the functionality that is contained inside of this other package named fmt

  - fmt is the name of a standard library package that is included with the go programming language by default

  - fmt is a shortened form of the word format

  - fmt library is used to print out a lot of different information specifically to the terminal, just to give us a better sense of debugging

  - main:
    - fmt (standard lib)
    - calculator (reusable package)
    - uploader (reusable package)

Note: Package fmt implements formatted I/O with functions analogous to C's printf and scanf

- What's that "func" thing?

  - func is short for functions

  <pre>
        <code>
            func main() {
                // function body
            }
        </code>
  </pre>

  - func: tells go that we are about to declare a function
  - main: sets the name of the function
  - (): list of arguments to pass the function
  - function body: call function runs this code

- How is the main.go file organized?
  - package main
    <br/>
    package declaration
  - import "fmt"
    <br/>
    import other packages that we need
  - func main() {...}
    <br/>
    declare functions - tell Go to do things

# Project

Cards:

- newDeck
  <br/>
  create a list of playing cards, essentially an array of strings

- print
  <br/>
  log out the contents of a deck of cards

- shuffle
  <br/>
  shuffles all the cards in a deck

- deal
  <br/>
  create a "hand" of cards

- saveToFile
  <br/>
  save a list of cards to a file on the local machine

- newDeckFromFile
  <br/>
  load a list of cards from the local machine

# Variables

- var card string = "Ace of Spades"

  - var: we are about to create a new variable
  - card: variable name
  - string: only a "string" will ever be assigned to this variable
  - "Ace of Spades": assign the value "Ace of Spades" to this variable

- card := "Ace of Spades"

  - We are relying upon the Go compiler to just kind of figure out that card is supposed to contain a string
    <br/>
    (colon equals operator - indicate new variable)

- Dynamic Types:
  A dynamically typed language is one in which the developers essentially do not care what values we are assigning to any given variable.
  <br/>
  E.g. Javascript, Ruby, Python

- Static Types:
  A variable which have the property to preserve its value from its previous scope.
  <br/>
  E.g. C++, Java, Go

- Basic Go Types:
  - bool (true/false)
  - string ("Hi!"/"Hows it going?")
  - int (0/-10000/99999)
  - float64 (10.0000001/0.000009/-100.003)

Note: Blank identifier - ignore unused variable "\_"

# Functions and Return Types

Go is always going to expect us to label the type of data that is being exchanged around our different functions inside of our program.

- Return Value
  <pre>
        <code>
            func newCard() string {
                // function body
            }
        </code>
  </pre>

  - newCard(): define a function called "newCard"
  - string: when executed, this function will return a value of type "string"

- Multiple Return Values
  <pre>
        <code>
              func deal(d deck, handSize int) (deck, deck) {
                // function body
              }
        </code>
  </pre>

  - (d deck, handSize int): receiver
  - (deck, deck): return 2 separate values of type deck

# Slices and For Loops

Array: Fixed length list of things
<br/>
Slice: An array that can grow or shrink

- Both Array and Slices must be defined with a datatype.

- Every element in a slice must be of same type.
  E.g. "Five of Spades", "Three of Diamonds", "Five of Diamonds" (Correct)
  "Five of Spades", 55525235, "Five of Diamonds" (Incorrect)

- To initialize Slice
  cards := []string{"Ace of Diamonds", newCard()}

- To add new element to Slice
  cards = append(cards, "Six of Spades")

- Slices are zero-indexed.

- Byte Slices

  - package ioutil implements some I/O utility functions
    - func WriteFile (filename string, data []byte, perm os.FileMode) error
      WriteFile writes data to a file named by filename. If the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing.

- For Loop
  <pre>
        <code>
            for index, card := range cards {
                fmt.Println(card)
            }
        </code>
  </pre>

  - index: index of this element in the array (array index)
  - card: current card we are iterating over (array element)
  - range cards: take the slice of "cards" and loop over it (array)
  - fmt.Println(card): run this one time for each card in the slice (print array elements)

# OO Approach vs Go Approach

- Object Oriented (OO) Approach:

  - Deck Class ---> Deck Instance
    <br/>

    <img width="250" height="300" alt="image" src="https://user-images.githubusercontent.com/102248503/188568125-9d9e21a3-54dd-4dc9-af5c-19cd6a956b39.png">

    <br/>

- Go Approach:

  - Base Go Types (string, integer, float, array, map)

    - we want to "extend" a base type and add some extra functionality to it

  - type deck []string

    - Tell Go we want to create an array of strings and add a bunch of functions specifically made to work with it
    - type gives us the ability to create a bunch of custom functions that only work with that deck type

  - Functions with 'deck' as a 'receiver'

    - A function with a receiver is like a "method" - a function that belongs to an "instance"

  - 'cards' folder:

    - main.go (code to create and manipulate a deck)
    - deck.go (code that describes what a deck is and how it works)
    - deck_test.go (code to automatically test the deck)
