package main

import (
    . "fmt"
)

var (
    p []int
    used []bool
    words map[byte]int
    finish bool
    S1 string
    S2 string
    S3 string
)

func dfs(start int, depth int) {
    if finish {
        return
    }

    p[depth] = start
    if depth == (len(words)-1) {
        check(p)
        return
    }
    used[start] = true
    for next:=0; next<10; next++ {
        if used[next] {
            continue
        }
        dfs(next, depth+1)
    }
    used[start] = false
}

func num(s string, p []int) int {
    n := 0
    for i:=0; i<len(s); i++ {
        if i == 0 && p[words[s[i]]] == 0 {
            return 0
        }
        _pow := 1
        for j:=0; j<(len(s)-i-1); j++ {
            _pow *= 10
        }
        n += p[words[s[i]]] * _pow
    }
    return n
}

func check(p []int) {
    n1 := num(S1, p)
    n2 := num(S2, p)
    n3 := num(S3, p)
    if n1 == 0 || n2 == 0 || n3 == 0 {
        return
    }
    if (n1 + n2) == n3 {
        Println(n1)
        Println(n2)
        Println(n3)
        finish = true
    }
    return
}

func main() {
    Scan(&S1)
    Scan(&S2)
    Scan(&S3)

    // 桁での判定、lenはbyteでいい
    // if (len(S3) - len(S1)) > 2 && (len(S3) - len(S2)) > 2 {
        // Println("UNSOLVABLE")
        // return
    // }

    // 単語map作る
    words = make(map[byte]int)
	for i := 0; i < len(S1); i++ {
		words[S1[i]] = 0
	}
	for i := 0; i < len(S2); i++ {
		words[S2[i]] = 0
	}
	for i := 0; i < len(S3); i++ {
		words[S3[i]] = 0
	}
    v := 0
    for k, _ := range words {
        words[k] = v
        v += 1
    }
    if len(words) > 10 {
        Println("UNSOLVABLE")
        return
    }

    // 初期化
    p = make([]int, len(words))
    used = make([]bool, 10)
    finish = false

    for start:=0; start<10; start++ {
        dfs(start, 0)
    }
    if !finish {
        Println("UNSOLVABLE")
    }
}
