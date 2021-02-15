package arra8

import "fmt"

func main() {
	var n, m, count, i, j, ans int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	b := make([]int, m)
	for i := range b {
		fmt.Scan(&b[i])
	}
	ans = 100000
	for ; i < m; i++ {
		a[b[i]-1]++
		if a[b[i]-1] == 1 {
			count++
		}
		if count == n {
			if ans > i-j {
				ans = i - j + 1
			}
			for ; j <= i; j++ {
				a[b[j]-1]--
				if a[b[j]-1] == 0 {
					count--
					j++
					break
				}
				if count == n {
					if ans > i-j {
						ans = i - j
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
