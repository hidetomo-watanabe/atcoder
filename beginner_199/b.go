package main

import (
    . "fmt"
)

func main() {
    var (
        N int
        A int
        B int
        A_MAX int
        B_MIN int
    )
    Scan(&N)
    A_MAX = 1
    for i:=0; i<N; i++ {
        Scan(&A)
        if A_MAX < A {
            A_MAX = A
        }
    }
    B_MIN = 1000
    for i:=0; i<N; i++ {
        Scan(&B)
        if B_MIN > B {
            B_MIN = B
        }
    }
    ans := B_MIN - A_MAX + 1
    if ans < 0 {
        Println(0)
    } else {
        Println(ans)
    }
}
