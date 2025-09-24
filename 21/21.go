package main

//5 plats
// une note par plat

var plats = [5]string{"poulet", "boeuf", "poisson", "tofu", "salade"}
var notes = [5]int{15, 12, 18, 14, 10}

func main() {
	for i := 0; i < len(plats); i++ {
		println("Le plat de", plats[i], "a une note de", notes[i], "/20")
	}
}
