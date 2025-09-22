package main

var note int

func main() {
	if note < 0 || note > 20 {
		println("Erreur")
	}
	if note < 8 {
		println("Échec")
	}
	if note >= 8 && note < 12 {
		println("Passable")
	}
	if note >= 12 && note < 15 {
		println("Bien")
	}
	if note >= 15 && note <= 20 {
		println("Excellent")
	}
}
