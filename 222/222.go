package main

//ajouter toute les dépense a somme

var depenses = []int{10, 20, 30, 40, 50}

func main() {
	somme := 0
	for i := 0; i < len(depenses); i++ {
		somme += depenses[i]
	}
	println(somme)
}
