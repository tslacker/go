package hashpb40

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	a := make(map[string]string)
	ans := []string{}
	a[" "] = " "
	for i, j := range s {
		a[string(j)] = string(i + 'a')
	}
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	s = sc.Text()
	for _, j := range s {
		ans = append(ans, string(j))
	}
	for i := 0; i < n; i++ {
		for j := range ans {
			ans[j] = a[ans[j]]
		}
	}
	for _, i := range ans {
		fmt.Print(i)
	}
}
