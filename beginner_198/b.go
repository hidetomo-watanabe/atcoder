package main

import (
    "fmt"
    "strings"
)

func main() {
    var N string
    fmt.Scan(&N)
    org := strings.Split(N, "")
    zero_num := 0
    reversed := []string{}
    for i:=0; i<len(org); i++ {
        var _n = org[len(org)-i-1]
        if len(reversed) == 0 && _n == "0" {
            zero_num += 1
        } else {
            reversed = append(reversed, _n)
        }
    }
    if strings.Join(org[:len(org)-zero_num], "") == strings.Join(reversed, "") {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}
