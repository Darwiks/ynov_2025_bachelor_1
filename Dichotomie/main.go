package main

import (
	"fmt"
	"time"
)

var tab = []int{1, 12, 34, 56, 76, 82, 85, 92, 109, 154, 189, 200, 234, 256, 300, 356, 400, 450, 500, 600, 700, 800, 900, 1000, 1200, 1500, 1800, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 6000, 7000, 8000, 9000, 10000, 12000, 15000, 18000, 20000, 25000, 30000, 35000, 40000, 45000, 50000, 60000, 70000, 80000, 90000, 100000}

func dichotomie(tab []int, x int) bool {
	debut := 0
	fin := len(tab) - 1
	for debut <= fin {
		m := (debut + fin) / 2
		if x == tab[m] {
			return true
		}
		if x > tab[m] {
			debut = m + 1
		}
		if x < tab[m] {
			fin = m - 1
		}
	}
	return false
}

func main() {
	fmt.Printf("Entrez un nombre à chercher : ")
	var x int
	fmt.Scan(&x)

	start := time.Now()
	result := dichotomie(tab, x)
	elapsed := time.Since(start)

	println(result)
	fmt.Printf("Temps écoulé : %s\n", elapsed)
}
