package main

import (
    "fmt"
    "math"
)

func main() {
    var N int
    fmt.Scan(&N)
    ans := -1
    for i := 1; i <= int(math.Sqrt(float64(N))); i++ {
        d := N / i
        m := N % i
        if m == 0 {
            ans = d + i - 2
        }
    }
    fmt.Println(ans)
}
