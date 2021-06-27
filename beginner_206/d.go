package main

import (
    . "fmt"
    "os"
    "bufio"
    "strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	x, _ := strconv.Atoi(next())
	return x
}

func chars(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}

type UnionFind struct {
    par []int
}

/* コンストラクタ */
func newUnionFind(N int) *UnionFind {
    u := new(UnionFind)
    u.par = make([]int, N)
    for i := range u.par {
        u.par[i] = -1
    }
    return u
}

/* xの所属するグループを返す */
func (u UnionFind) root(x int) int {
    if u.par[x] < 0 {
        return x
    }
    u.par[x] = u.root(u.par[x])
    return u.par[x]
}

/* xの所属するグループ と yの所属するグループ を合体する */
func (u UnionFind) unite(x, y int) {
    xr := u.root(x)
    yr := u.root(y)
    if xr == yr {
        return
    }
    u.par[yr] += u.par[xr]
    u.par[xr] = yr
}

/* xとyが同じグループに所属するかどうかを返す */
func (u UnionFind) same(x, y int) bool {
    if u.root(x) == u.root(y) {
        return true
    }
    return false
}

/* xの所属するグループの木の大きさを返す */
func (u UnionFind) size(x int) int {
    return -u.par[u.root(x)]
}

func main() {
    // init
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()

    var (
        N int
        A int
        ans int
    )
    N = nextInt()

    s1 := []int {}
    s2 := []int {}
    if N % 2 == 0 {
        for i:=0; i<N/2; i++ {
            A = nextInt()
            s1 = append(s1, A)
            s2 = append(s2, 0)
        }
        for i:=0; i<N/2; i++ {
            A = nextInt()
            s2[N/2-1-i] = A
        }
    } else {
        for i:=0; i<(N-1)/2; i++ {
            A = nextInt()
            s1 = append(s1, A)
            s2 = append(s2, 0)
        }
        A = nextInt()
        for i:=0; i<(N-1)/2; i++ {
            A = nextInt()
            s2[(N-1)/2-1-i] = A
        }
    }

    max_a := 200001
    u := newUnionFind(max_a)
    for i:=0; i<len(s1); i++ {
        if s1[i] == s2[i] {
            continue
        }

        u.unite(s1[i], s2[i])
    }

    ans = 0
    for i:=0; i<max_a; i++ {
        if u.size(i) > 1 && u.root(i) == i {
            ans += u.size(i) - 1
        }
    }
    Fprintln(out, ans)
}
