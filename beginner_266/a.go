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
        S string
        ans string
    )
    S = next()

    i := (len(S) - 1) / 2
    ans = S[i: i+1]

    Fprintln(out, ans)
}
