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

func nextFloat() float64 {
	x, _ := strconv.ParseFloat(next(), 64)
	return x
}

func chars(s string) []string {
	ret := make([]string, len([]rune(s)))
	for i, v := range []rune(s) {
		ret[i] = string(v)
	}
	return ret
}

var ans int
func dfs(N int, X int, a [][]int, pos int, val int) {
    if pos == N {
        if val == X {
            ans += 1
        }
        return
    }
    for _, _a := range a[pos] {
        if val > X {
            continue
        }
        dfs(N, X, a, pos+1, val*_a)
    }
}

func main() {
    // init
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()

    var (
        N int
        X int
    )
    N = nextInt()
    X = nextInt()

    a := make([][]int, N)
    for i:=0; i<N; i++ {
        l := nextInt()
        for j:=0; j<l; j++ {
            a[i] = append(a[i], nextInt())
        }
    }

    dfs(N, X, a, 0, 1)
    Fprintln(out, ans)
}
