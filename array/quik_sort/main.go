package main

import "fmt"

func quikSort(a []int, first int, last int) {
	pivot := (last + first) / 2
	l := first
	r := last
	for l <= r {
		for a[l] < a[pivot] {
			l++
		}
		for a[r] > a[pivot] {
			r--
		}
		if l <= r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	if first < r {
		quikSort(a, first, r)
	}
	if l < last {
		quikSort(a, l, last)
	}
}

func sort(a []int) {
	quikSort(a, 0, len(a)-1)
}

func main() {
	a := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Println(a)

	sort(a)
	fmt.Println(a)
}
