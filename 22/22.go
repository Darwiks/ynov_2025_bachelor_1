package main

var places = [25]bool{true, true, true, true, true, false, true, true, true, true, false, true, true, true, true, true, true, true, true, true, false, false, true, true, true}

func main() {
	count := 0
	for i := 0; i < len(places); i++ {
		if places[i] == true {
			println("Place", i+1, "est libre")
			count++
		} else {
			println("Place", i+1, "est occupée")
		}
	}
	println("Nombre total de places libres :", count)
}
