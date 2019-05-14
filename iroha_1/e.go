package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func ri() (n int) {
	sc.Scan()
	for _, v := range sc.Bytes() {
		n = n*10 + int(v-48)
	}
	return
}

func main() {
	// init
	sc.Split(bufio.ScanWords)
	N, A, B := ri(), ri(), ri()
	var D_LIST []int
	if B > 0 {
		var num int
		for i := 0; i < B; i++ {
			num = ri()
			D_LIST = append(D_LIST, num)
		}
	}

	// loop
	sort.Sort(sort.IntSlice(D_LIST))
	D_LIST = append([]int{0}, D_LIST...)
	// last needs not date
	D_LIST = append(D_LIST, N+1)
	date := B
	for i := 1; i < len(D_LIST); i++ {
		date += (D_LIST[i] - D_LIST[i-1] - 1) / A
	}

	// output
	fmt.Println(N - date)
}
