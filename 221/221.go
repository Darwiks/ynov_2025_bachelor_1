package main

// pour toute la classe 
// afficher le prénom des collègues 

var classe = []string{"Alice", "Bob", "Charlie", "Diana"}

func main() {
	for i := 0; i < len(classe); i++ {
		println(classe[i])
	}
}
