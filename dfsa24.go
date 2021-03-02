package main

import "fmt"

var ans int
var n, k int

func main() {
	fmt.Scan(&n, &k)
	if n > k {
		fmt.Println(0)
		return
	}
	isPermList([]int{})
	fmt.Println(ans)
}

func isPermList(head []int) {
	if len(head) == n {
		var b []int
		for i := 1; i <= n*2; i++ {
			var check bool
			for j := 0; j < n; j++ {
				if i == head[j] {
					check = true
					break
				}
			}
			if !check {
				b = append(b, i)
			}
		}
		sum := 0
		for i := range head {
			if head[i]-b[i] <= 0 {
				sum += b[i] - head[i]
			} else {
				sum += head[i] - b[i]
			}
		}
		if sum <= k {
			ans++
		}
		return
	} else {
		for v := 1; v <= 2*n; v++ {
			if len(head) == 0 || head[len(head)-1] < v {
				headx := append(head, v)
				isPermList(headx)
			}
		}
	}
}
