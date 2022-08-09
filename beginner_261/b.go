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

func checkA(N int, A [][]string) string {
    ans := "correct"
    _a := ""
    for i:=0; i<N; i++ {
        if A[i][i] != "-" {
            ans = "incorrect"
            return ans
        }
        for j:=i+1; j<N; j++ {
            _a = A[i][j] + A[j][i]
            if !(_a == "WL" || _a == "LW" || _a == "DD") {
                ans = "incorrect"
                return ans
            } else {
            }
        }
    }
    return ans
}

func main() {
    // init
    buf := make([]byte, 1024*1024)
    sc.Buffer(buf, bufio.MaxScanTokenSize)
    sc.Split(bufio.ScanWords)
    defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()

    var (
        N int
        A [][]string
        ans string
    )
    N = nextInt()
    for i:=0; i<N; i++ {
        A = append(A, []string{})
        A[i] = strings.Split(next(), "")
    }

    ans = checkA(N, A)
    Fprintln(out, ans)
}
