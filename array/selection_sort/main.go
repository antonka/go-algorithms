package main

import "fmt"

func sort(a []int) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	a := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println(a)

	sort(a)
	fmt.Println(a)
}
