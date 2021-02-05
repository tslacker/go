package main

import (
	"fmt"
	"sort"
)

type t struct {
	x, y int
}

type sans struct {
	men, hen int
}

func pa21() {
	var h, w int
	var get string
	fmt.Scan(&h, &w)
	sa := make([][]string, h+2)
	sb := make([][]bool, h+2)
	ans := []sans{}
	for i := range sa {
		sa[i] = make([]string, w+2)
		sb[i] = make([]bool, w+2)
		if i > 0 {
			fmt.Scan(&get)
		}
		for j := 0; j <= w+1; j++ {
			if i == 0 || i == h+1 || j == 0 || j == w+1 {
				sa[i][j] = "."
			} else {
				sa[i][j] = string(get[j-1])
				if get[j-1] == '#' {
					sb[i][j] = true
				}
			}
		}
	}
	vy := []int{1, 0, -1, 0}
	vx := []int{0, 1, 0, -1}
	for i := 0; i <= h; i++ {
		for j := 0; j <= w; j++ {
			if sa[i][j] == "#" {
				hen := 0
				men := 0
				now := []t{t{y: i, x: j}}
				for len(now) > 0 {
					next := []t{}
					for k := 0; k < len(now); k++ {
						if sa[now[k].y][now[k].x] == "#" {
							hen += 4
							men += 1
						}
						for l := 0; l < len(vy); l++ {
							ny := now[k].y + vy[l]
							nx := now[k].x + vx[l]
							if sb[ny][nx] == true && sa[now[k].y][now[k].x] == "#" {
								hen--
								if sa[ny][nx] == "#" {
									next = append(next, t{y: ny, x: nx})
								}
							}
						}
						sa[now[k].y][now[k].x] = "."
					}
					now = next
				}
				ans = append(ans, sans{men: men, hen: hen})
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].men == ans[j].men {
			return ans[i].hen > ans[j].hen
		} else {
			return ans[i].men > ans[j].men
		}
	})
	for i := range ans {
		fmt.Println(ans[i].men, ans[i].hen)
	}
}
