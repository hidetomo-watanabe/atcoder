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
        K int
        A string
        B string
        ans int
    )
    K = nextInt()
    A = next()
    B = next()

    A2 := 0
    for i:=0; i<len(A); i++ {
        n, _ := strconv.Atoi(string(A[i:i+1]))
        A2 += n * int(math.Pow(float64(K), float64(len(A)-i-1)))
    }
    B2 := 0
    for i:=0; i<len(B); i++ {
        n, _ := strconv.Atoi(string(B[i:i+1]))
        B2 += n * int(math.Pow(float64(K), float64(len(B)-i-1)))
    }

    ans = A2 * B2
    Fprintln(out, ans)
}
