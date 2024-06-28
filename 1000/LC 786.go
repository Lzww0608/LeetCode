func kthSmallestPrimeFraction(arr []int, k int) []int {
    var (
        eps float64 = 1e-8
        l float64 = 0
        r float64 = 1
    )
    n := len(arr)
    a, b := 0, 0
    check := func(x float64) int {
        ans := 0
        for i, j := 0, 1; j < n; j++ {
            for float64(arr[i+1]) / float64(arr[j]) <= x {
                i++
            }
            if float64(arr[i]) / float64(arr[j]) <= x {
                ans += i + 1
            }
            if abs(float64(arr[i]) / float64(arr[j]) - float64(x)) < eps {
                a, b = arr[i], arr[j]
            }

        }
        return ans
    }

    for r - l > eps {
        mid := l + ((r - l) / 2.0)
        if check(mid) >= k {
            r = mid
        } else {
            l = mid
        }
    }

    return []int{a, b}
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }

    return x
}
