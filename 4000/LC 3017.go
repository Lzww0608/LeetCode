func countOfPairs(n int, x int, y int) []int64 {
    if x > y {
        x, y = y, x 
    }

    ans := make([]int64, n + 1)
    for i := 1; i <= n; i++ {
        // 直接走，不走x-y
        if abs(i - x) + 1 > abs(i - y) {
            ans[1] += 2
            ans[n - i + 1] -= 2
        } else {
            d := abs(i - x) + 1
            // d + y - sep == sep - i 
            sep := (y + i + d) / 2

            ans[1] += 2
            ans[sep - i + 1] -= 2

            ans[d] += 2
            ans[d + y - sep] -= 2

            ans[d + 1] += 2
            ans[d + n - y + 1] -= 2
        }
    }

    for i := 1; i <= n; i++ {
        ans[i] += ans[i-1]
    }

    return ans[1:]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}