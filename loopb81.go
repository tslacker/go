package loopb81

import "fmt"

func main() {
	var h, w, ans int
	x := []int{0, 1, 0, -1}
	y := []int{1, 0, -1, 0}
	fmt.Scan(&h, &w)
	s := make([]string, h+2)
	for i := 0; i <= h+1; i++ {
		if i == 0 || i == h+1 {
			for j := 0; j <= w+1; j++ {
				s[i] += "."
			}
		} else {
			fmt.Scan(&s[i])
			s[i] = "." + s[i] + "."
		}
	}
	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if s[i][j] == '#' {
				ans += 4
				for k := 0; k < 4; k++ {
					if s[i+x[k]][j+y[k]] == '#' {
						ans--
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
