package main

import "fmt"

func sort(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
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
