package main

import (
	"bufio"
	"fmt"
	"os"
)

var d, n, pos, max, min int
var a []int
var c bool
var test [1001][2002]bool

func main() {
	var sc = bufio.NewReader(os.Stdin)
	fmt.Fscan(sc, &d, &n)
	max = d
	min = d * -1
	var b int
	for i := 0; i < n; i++ {
		fmt.Fscan(sc, &b)
		a = append(a, b)
	}
	f(0, 0, "")
}

func f(pos, index int, ans string) {
	p := pos + 1000
	if test[index][p] {
		return
	}
	test[index][p] = true
	if c {
		return
	}
	if index == n {
		fmt.Println(ans)
		c = true
		return
	}
	if pos-a[index] >= min {
		f(pos-a[index], index+1, ans+"L")
	}
	if pos+a[index] <= max {
		f(pos+a[index], index+1, ans+"R")
	}
}
