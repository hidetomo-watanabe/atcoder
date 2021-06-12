package main

import (
    . "fmt"
    "os"
    "bufio"
    "sort"
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

func main() {
    // init
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()

    var (
        N int
        K int
        A int
        B int
        ans int
    )
    N, K = nextInt(), nextInt()

    a2b := make(map[int]int)
    for i:=0; i<N; i++ {
        A, B = nextInt(), nextInt()
        _, ok := a2b[A]
        if ok {
            a2b[A] += B
        } else {
            a2b[A] = B
        }
    }
    as := []int{}
    for a, _ := range a2b {
        as = append(as, a)
    }
    sort.Ints(as)

    ans = 0
    m := K
    for _, a := range as {
        b, ok := a2b[a]
        if !ok {
            continue
        }

        if (a - ans) <= m {
            m -= (a - ans)
            m += b
            ans = a
        } else {
            ans += m
            m = 0
        }

        if m == 0 {
            break
        }
    }
    if m > 0 {
        ans += m
    }
    Fprintln(out, ans)
}
