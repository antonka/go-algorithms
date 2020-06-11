package main

import "fmt"

func shift(a []int, times int, withoutNull bool) {
	l := len(a)
	if times > 0 {
		for j := 0; j < times; j++ {
			tmp := 0
			for i := l - 1; i > 0; i-- {
				if tmp == 0 && withoutNull {
					tmp = a[i]
				}
				a[i] = a[i-1]
			}
			a[j] = tmp
		}
	} else {
		times = times * -1
		last := l - 1
		for j := last; j > last-times; j-- {
			tmp := 0
			for i := 0; i < last; i++ {
				if tmp == 0 && withoutNull {
					tmp = a[i]
				}
				a[i] = a[i+1]
			}
			a[j] = tmp
		}
	}
}

func main() {
	var a []int

	a = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)

	shift(a, 2, false)
	fmt.Println(a)

	a = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)

	shift(a, -2, true)
	fmt.Println(a)
}
