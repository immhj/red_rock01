package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Getnumber() (number int) {
	maxNum := 100
	time.Sleep(200)
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	return secretNumber
}

var q = []int{}

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
	n := 100
	for i = 0; i < 100; i++ {
		q = append(q, Getnumber())
	}
	my_sort(0, n-1, q)
	for i = 0; i < n; i++ {
		fmt.Println(q[i])
	}
}
