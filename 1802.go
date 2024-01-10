func maxValue(n int, index int, maxSum int) int {
    l, r := 1, maxSum
    cal := func(a, b int) int {
        if a > b {
            return (a + (a - b + 1)) * b / 2
        } else {
            return b - a + (a + 1) * a / 2
        }
    }
    for l < r {
        mid := l + ((r - l + 1) >> 1) 
        sum := cal(mid, index + 1) + cal(mid, n - index) - mid
        if sum <= maxSum {
            l = mid
        } else {
            r = mid - 1
        }
    }
    return l
}

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}

