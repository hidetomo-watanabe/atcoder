package main

import (
    . "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
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
        S string
        Q int
        T int
        A int
        B int
        r bool
    )
    N, S, Q = nextInt(), next(), nextInt()

    l := chars(S)
    r = false
    for i:=0; i<Q; i++ {
        T, A, B = nextInt(), nextInt(), nextInt()
        if T == 1 {
            a := A - 1
            b := B - 1
            if r {
                a = (a + N) % (2 * N)
                b = (b + N) % (2 * N)
            }
            l[a], l[b] = l[b], l[a]
        } else if T == 2 {
            r = !r
        }
    }
    if r {
        l = append(l[N:2*N], l[0:N]...)
    }
    Fprintln(out, strings.Join(l, ""))
}
