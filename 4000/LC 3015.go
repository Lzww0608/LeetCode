func countOfPairs(n int, x int, y int) []int {
    ans := make([]int, n)

    for i := 1; i < n; i++ {
        for j := i + 1; j <= n; j++ {
            d := min(j - i, abs(i - x) + 1 + abs(j - y), abs(i - y) + 1 + abs(j - x)) - 1
            ans[d] += 2
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}