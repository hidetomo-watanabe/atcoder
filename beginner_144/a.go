package main

import (
    "fmt"
)

func main() {
    var A int64
    var B int64
    fmt.Scan(&A)
    fmt.Scan(&B)
    if A <= 9 && B <= 9 {
        fmt.Println(A * B)
    } else {
        fmt.Println(-1)
    }
}
