package main

import (
	"fmt"
)

var q = []int{5, 1, 3, 4, 8}

func my_sort(l int, r int, q []int) {
	if l >= r {
		return
	}
	x := q[l]
	i := l - 1
	j := r + 1
	for i < j {
		i++
		for ; q[i] < x; i++ {

		}
		j--
		for ; q[j] > x; j-- {

		}
		if i < j {
			t := q[i]
			q[i] = q[j]
			q[j] = t
		}
	}
	my_sort(l, j, q)
	my_sort(j+1, r, q)
}
func main() {
	var i int
	var n int
	fmt.Scanf("%d", &n)
	my_sort(0, n-1, q)
	for i = 0; i < n; i++ {
		fmt.Println(q[i])
	}
}
