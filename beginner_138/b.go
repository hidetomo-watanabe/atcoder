package main

import (
	"bufio"
	"fmt"
    "os"
)

func main() {
	var N int
	fmt.Scan(&N)
	A_LIST := make([]float64, N)
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < N; i++ {
		fmt.Scan(&A_LIST[i])
	}
    ans := 0.0
	for i := 0; i < N; i++ {
        ans += 1 / A_LIST[i]
    }
    ans = 1.0 / ans
	fmt.Println(ans)
}
