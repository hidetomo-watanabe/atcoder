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
        ans string
    )
    N = nextInt()
    S = next()

    for i:=0; i<N; i++ {
        if S[i:i+1] == "1" {
            if i % 2 == 0 {
                ans = "Takahashi"
            } else {
                ans = "Aoki"
            }
            break
        }
    }

    Fprintln(out, ans)
}
