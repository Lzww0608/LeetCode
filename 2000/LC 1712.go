const MOD int = 1_000_000_007

func waysToSplit(nums []int) (ans int) {
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i+1] = pre[i] + x
    }

    for i := 1; i < n - 1; i++ {
        if (pre[i] * 3 > pre[n]) {
            break
        }

        l, r := i + 1, n
        for l < r {
            mid := l + ((r - l) >> 1)
            if pre[mid] - pre[i] < pre[i] {
                l = mid + 1
            } else {
                r = mid
            }
        }
        left := l

        l, r = l, n
        for l < r {
            mid := l + ((r - l) >> 1)
            if pre[n] - pre[mid] >= pre[mid] - pre[i] {
                l = mid + 1
            } else {
                r = mid 
            }
        }
        right := l

        ans += right - left
    }

    return max(0, ans % MOD)
}
