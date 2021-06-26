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
        N int
        ans int
    )
    N = nextInt()

    _ans := int(math.Ceil(math.Sqrt(2 * float64(N) - 1/4) - 1/2))

    // check
    _ans_list := []int{_ans - 1, _ans, _ans + 1}
    for _, _a := range _ans_list {
        _r := _a * (_a + 1) / 2
        if _r >= N {
            ans = _a
            break
        }
    }
    Fprintln(out, ans)
}
