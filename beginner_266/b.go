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
        MOD int
        N int
        ans int
    )
    MOD = 998244353
    N = nextInt()

    if (N % MOD == 0) {
        ans = 0
    } else if (N >= 0) {
        ans = N % MOD
    } else {
        ans = N % MOD + MOD
    }

    Fprintln(out, ans)
}
