package main

import "fmt"

func sort(a []int) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
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
