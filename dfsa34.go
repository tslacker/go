package dfsa34

import "fmt"

var n, x, sum, cnt, ans int
var a []int

func main() {
	fmt.Scan(&n, &x)
	a = make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	f(0, 0, 0)
	fmt.Println(ans)
}

func f(i, s, c int) {
	if i == n {
		if cnt < c {
			cnt = c
			ans = x - s
		} else if cnt == c {
			if ans > x-s {
				ans = x - s
			}
		}
		return
	}
	f(i+1, s, c)
	if x-s-a[i] >= 0 {
		f(i+1, s+a[i], c+1)
	}
	return
}
