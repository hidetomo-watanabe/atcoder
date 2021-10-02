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

func main() {
    // init
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()

    var (
        N int
        A_LIST []int
        A_SUM int
        X int
        ans int
    )
    N = nextInt()
    for i:=0; i<N; i++ {
        a := nextInt()
        A_LIST = append(A_LIST, a)
        A_SUM += a
    }

    X = nextInt()
    ans = N * (X / A_SUM)
    tmp_x := A_SUM * (ans / N)
    for i:=0; i<N; i++ {
        tmp_x += A_LIST[i]
        ans += 1
        if tmp_x > X {
            break
        }
    }

    Fprintln(out, ans)
}
