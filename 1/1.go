package main

import "fmt"

var majeure bool
var age int
var prenom string

func main() {
	if age >= 18 {
		majeure = true
	} else {
		majeure = false
	}
	if majeure == true {
		fmt.Println(prenom + " Envoyer le document officiel")
	} else {
		return
	}
}
