package main

import (
    . "fmt"
)

func main() {
    var (
        A int64
        B int64
        C int64
    )
    Scan(&A)
    Scan(&B)
    Scan(&C)
    if A * A + B * B < C * C {
        Println("Yes")
    } else {
        Println("No")
    }
}
