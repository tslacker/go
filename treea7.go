package treea7

import (
	"bufio"
	"fmt"
	"os"
)

type UnionFind struct {
	par []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

func (u UnionFind) root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

func (u UnionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func main() {
	var n, m, t, a, b int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	// インスタンスの生成
	u := newUnionFind(n)
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &t, &a, &b)
		a--
		b--
		if t == 0 {
			// 各メソッドの利用
			u.unite(a, b)
		} else {
			if u.same(a, b) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
