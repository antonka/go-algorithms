package main

import "fmt"

func revers(a []int) {
	l := len(a) - 1
	i := 0
	for i < l {
		a[i], a[l] = a[l], a[i]
		i++
		l--
	}
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(a)

	revers(a)
	fmt.Println(a)
}
