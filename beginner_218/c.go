package main

import (
    . "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
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
        ss [][2]int
        ts0 [][2]int
        ts90 [][2]int
        ts180 [][2]int
        ts270 [][2]int
        ans string
    )

    N = nextInt()

    for i:=0; i<N; i++ {
        for j, _s := range next() {
            if string(_s) == "#" {
                _xy := [2]int{i, j}
                ss = append(ss, _xy)
            }
        }
    }

    // left rotation
    for i:=0; i<N; i++ {
        for j, _t := range next() {
            for _, _rot := range [4]int{0, 90, 180, 270} {
                var _x int
                var _y int
                if string(_t) == "#" {
                    _rotate_base := 2 * (N/2)
                    if _rot == 0 {
                        _x = i
                        _y = j
                        ts0 = append(ts0, [2]int{_x, _y})
                    } else if _rot == 90 {
                        _x = _rotate_base - j
                        _y = i
                        ts90 = append(ts90, [2]int{_x, _y})
                    } else if _rot == 180 {
                        _x = _rotate_base - i
                        _y = _rotate_base - j
                        ts180 = append(ts180, [2]int{_x, _y})
                    } else if _rot == 270 {
                        _x = j
                        _y = _rotate_base - i
                        ts270 = append(ts270, [2]int{_x, _y})
                    }
                }
            }
        }
    }

    if len(ss) != len(ts0) {
        Fprintln(out, "No")
        return
    }

    for i, _ts := range [4][][2]int{ts0, ts90, ts180, ts270} {
        if i != 0 {
            sort.Slice(_ts, func(j, k int) bool {
                if _ts[j][0] < _ts[k][0] {
                    return true
                } else if _ts[j][0] == _ts[k][0] {
                    if _ts[j][1] < _ts[k][1] {
                        return true
                    }
                }
                return false
            })
        }

        ans = "Yes"
        _parallel := [2]int{ss[0][0] - _ts[0][0], ss[0][1] - _ts[0][1]}
        for j, _t := range _ts {
            if (ss[j][0] - _t[0]) != _parallel[0] || (ss[j][1] - _t[1] != _parallel[1]) {
                ans = "No"
                break
            }
        }

        if ans == "Yes" {
            Fprintln(out, ans)
            break
        }
    }

    if ans == "No" {
        Fprintln(out, ans)
    }

    return
}
