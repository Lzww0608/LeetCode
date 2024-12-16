const N int = 11
func getProbability(balls []int) float64 {
    factor := [N]int{}
    factor[0] = 1
    for i := 1; i < N; i++ {
        factor[i] = factor[i-1] * i
    }

    n := len(balls)
    sum := 0
    for _, x := range balls {
        sum += x 
    }
    var q float64 = 1.0
    for i := 1; i <= sum / 2; i++ {
        q *= float64(i + sum / 2) / float64(i) / 4.0 
    }

    left := [N]int{}
    left[n-1] = balls[n-1]
    for i := n - 2; i >= 0; i-- {
        left[i] = left[i+1] + balls[i]
    }

    combination := func(a, b int) int {
        return factor[a] / (factor[b] * factor[a-b])
    }

    var dfs func(int, int, int) float64 
    dfs = func(m, sum, color int) float64 {
        if m == n {
            if sum == 0 && color == 0 {
                return 1
            }
            return 0
        }

        if abs(sum) > left[m] {
            return 0
        }
        var res float64 = 0
        for i := 0; i <= balls[m]; i++ {
            col := 0
            if i == 0 {
                col = -1
            } else if i == balls[m] {
                col = 1
            }
            res += float64((combination(balls[m], i))) / float64(int32(1) << balls[m]) * dfs(m + 1, sum + i - (balls[m] - i), color + col)
        }

        return res
    }

    return dfs(0, 0, 0) / q
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

