func minCost(a []int, b []int, k int64) int64 {
    n := len(a)
    sum := 0
    for i := range n {
        sum += abs(a[i] - b[i])
    }
    ans := sum
    if ans < int(k) {
        return int64(ans)
    }

    sort.Ints(a)
    sort.Ints(b)
    sum = int(k)
    for i := range n {
        sum += abs(a[i] - b[i])
    }
    ans = min(ans, sum)
    return int64(ans)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}