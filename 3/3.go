package main

import "fmt"

var prix float64
var donne float64

func main() {
	if donne < prix {
		fmt.Println("Il manque de l'argent")
	} else if donne > prix {
		fmt.Println("Rendre la monnaie")
	} else {
		fmt.Println("Merci pour le paiement exact")
	}
}
