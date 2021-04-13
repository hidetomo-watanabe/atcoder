package main

import (
    . "fmt"
    "math"
)

func main() {
    var R int
    var X int
    var Y int
    Scan(&R)
    Scan(&X)
    Scan(&Y)

    // 誤差対策でfloat64変換は必要な時に行う
    var goal_r2 = X * X + Y * Y
    var ans int
    // ゴールがRよりも小さいとき
    if goal_r2 < (R * R) {
        ans = 2
    } else {
        ans = int(math.Sqrt(float64(goal_r2)) / float64(R))
        // ans補正
        if (R * R * ans * ans) != goal_r2 {
            ans += 1
        }
    }
    Println(ans)
}
