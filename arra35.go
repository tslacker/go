package main

import (
	"fmt"
	"sort"
)

var ans []int
var memo [101][101]bool
var n int

func main() {
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	f(0, 0, a)
	sort.Ints(ans)
	fmt.Println(len(ans))
	for _, i := range ans {
		fmt.Println(i)
	}
}

func f(sum, index int, a []int) {
	if memo[sum][index] {
		return
	}
	memo[sum][index] = true
	if index == n {
		ans = append(ans, sum)
		return
	}
	f(sum+a[index], index+1, a)
	f(sum, index+1, a)
}
