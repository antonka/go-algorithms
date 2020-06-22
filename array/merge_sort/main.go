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

func sort(a []int) []int {

	l := len(a)
	if l == 1 {
		return a
	}

	m := l / 2
	b := sort(a[:m])
	c := sort(a[m:])

	return merge(b, c)
}

func main() {
	a := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println(a)

	a = sort(a)
	fmt.Println(a)
}
