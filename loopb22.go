package loopb22

import "fmt"

func main() {
	var m, n, k, a, max int
	fmt.Scan(&m, &n, &k)
	ans := make([]int, m+1)
	for i := 0; i < k; i++ {
		fmt.Scan(&a)
		if n > 0 {
			n--
			ans[a]++
		}
		for j := 1; j <= m; j++ {
			if j != a {
				if ans[j] > 0 {
					ans[j]--
					ans[a]++
				}
			}
		}
	}
	for _, i := range ans {
		if max < i {
			max = i
		}
	}
	for i, j := range ans {
		if max == j {
			fmt.Println(i)
		}
	}
}
