package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func ri() (n int) {
	sc.Scan()
	for _, v := range sc.Bytes() {
		n = n*10 + int(v-48)
	}
	return
}

func primeFactors(n int) []int {
	factors := make([]int, 0)
	i := 2
	for i*i <= n {
		r := n % i
		if r != 0 {
			i += 1
		} else if r == 0 {
			n /= i
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func main() {
	// init
	sc.Split(bufio.ScanWords)
	N, K := ri(), ri()

	// factoring
	factors := primeFactors(N)
	// ans
	var ans []int
	l := len(factors)
	if K > l {
		ans = append(ans, -1)
	} else if K == l {
		ans = factors
	} else if K < l {
		ans = factors
		for i := 0; i < (l - K); i++ {
			ans = append(ans[:len(ans)-2], ans[len(ans)-1]*ans[len(ans)-2])
		}
	}

	// output
	str := ""
	for _, v := range ans {
		str += fmt.Sprintf("%d ", v)
	}
	str = strings.TrimRight(str, " ")
	fmt.Println(str)
}
