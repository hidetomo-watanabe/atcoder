package main

import (
    . "fmt"
    "os"
    "bufio"
    "strconv"
    "math"
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
        L1 float64
        R1 float64
        L2 float64
        R2 float64
        ans int
    )
    L1 = nextFloat()
    R1 = nextFloat()
    L2 = nextFloat()
    R2 = nextFloat()

    ans = int(math.Min(R1, R2) - math.Max(L1, L2))
    if ans < 0 {
        ans = 0
    }

    Fprintln(out, ans)
}
