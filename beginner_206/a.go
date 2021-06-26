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
        N float64
        ans string
    )
    N = float64(nextInt())

    _n := math.Floor(1.08 * N)
    if _n > 206 {
        ans = ":("
    } else if _n == 206 {
        ans = "so-so"
    } else {
        ans = "Yay!"
    }
    Fprintln(out, ans)
}
