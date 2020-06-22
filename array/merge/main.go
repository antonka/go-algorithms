package main

import "fmt"

func merge(a []int, b []int) []int {

	la := len(a)
	lb := len(b)

	r := make([]int, la+lb)
	i, j, x := 0, 0, 0

	for i < la && j < lb {
		if a[i] < b[j] {
			r[x] = a[i]
			x++
			i++
		} else {
			r[x] = b[j]
			x++
			j++
		}
	}

	for i < la {
		r[x] = a[i]
		x++
		i++
	}

	for j < lb {
		r[x] = b[j]
		x++
		j++
	}

	return r
}

func main() {
	a := []int{1, 4, 5, 7, 8, 80}
	b := []int{2, 3, 9, 10, 12, 20, 34}
	c := merge(a, b)
	fmt.Println(c)
}
