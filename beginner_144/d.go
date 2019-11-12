package main

import (
    "fmt"
    "math"
)

func main() {
    var A, B, X float64
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&X)
    var ans float64
    if X <= A * A * B / 2 {
        ans = math.Atan(A * B * B / (2 * X))
    } else {
        ans = math.Atan(2 * (A * A * B - X) / (A * A * A))
    }
    ans = ans * 180 / math.Pi
    fmt.Println(ans)
}
