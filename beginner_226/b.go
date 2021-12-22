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

func _contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func _eqArrs(a []int, b []int) bool {
    // if len(a) != len(b) {
    if a[0] != b[0] {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func main() {
    // init
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	// sc.Split(bufio.ScanWords)
	defer out.Flush() // !!!!caution!!!! you must use Fprint(out, ) not Print()

    var (
        N int
        ans int
    )
    N = nextInt()
    ans = 0

    // setで追加
    set := make(map[string]struct{})
    for i:=0; i<N; i++ {
        // 配列のstr作成
        arr := next()
        set[arr] = struct{}{}
    }

    ans = len(set)
    Fprintln(out, ans)
}
