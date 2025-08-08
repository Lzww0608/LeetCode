func maxScore(prices []int) int64 {
    m := make(map[int]int)
    for i, x := range prices {
        m[x - i] += x
    }

    ans := 0
    for _, v := range m {
        ans = max(ans, v)
    }

    return int64(ans)
}