package arra28

import "fmt"

func main() {
	var s, p, ans int
	fmt.Scan(&s, &p)
	a := make(map[int]int)
	a[0] = s
	for i := 0; i <= p; i++ {
		for j := 0; j+i <= p; j++ {
			a[i+j] = Max(a[i+j], a[i]*(100+j)/100)
		}
	}
	for _, i := range a {
		ans = Max(ans, i)
	}
	fmt.Println(ans)
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
