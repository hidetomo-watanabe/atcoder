package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    N := nextInt()
    Q := nextInt()

    trees := make([][]int, N)
    for i := 0; i < N-1; i++ {
        a := nextInt()
        b := nextInt()
        trees[a-1] = append(trees[a-1], b-1)
        trees[b-1] = append(trees[b-1], a-1)
    }

    ans := make([]int, N)
    for i := 0; i < Q; i++ {
        p := nextInt()
        x := nextInt()
        ans[p-1] += x
    }

    ans = updateAns(trees, -1, 0, ans)

    for i := 0; i < N-1; i++ {
        fmt.Printf("%d ", ans[i])
    }
    fmt.Println(ans[N-1])
}

func next() string{
    sc.Scan()
    return sc.Text()
}

func nextInt() int{
    a, _ := strconv.Atoi(next())
    return a
}

func updateAns(trees [][]int, org int, parent int, ans[]int) []int {
    var child int
    for i := 0; i < len(trees[parent]); i++ {
        child = trees[parent][i]
        if child == org {
            continue
        }
        ans[child] += ans[parent]
        updateAns(trees, parent, child, ans)
    }
    return ans
}
