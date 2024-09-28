const MOD int = 1_000_000_007
func rangeSum(nums []int, n int, left int, right int) int {
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i+1] = pre[i] + x
    }
    ppre := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        ppre[i] = ppre[i-1] + pre[i]
    }

    getSumCnt := func(x int) (int, int) {
        sum, cnt := 0, 0
        for i, j := 0, 0; i < n; i++ {
            for j < n && pre[j+1] - pre[i] < x {
                j++
            }
            cnt += j - i
            sum += ppre[j] - ppre[i] - (j - i) * pre[i]
        }

        return sum, cnt
    }

    getKthSum := func(k int) int {
        l, r := 0, pre[n]
        for l < r {
            mid := r - ((r - l) >> 1)
            _, cnt := getSumCnt(mid)
            if cnt < k {
                l = mid
            } else {
                r = mid - 1
            }
        }

        return l
    }
    
    f := func(x int) int {
        t := getKthSum(x)
        sum, cnt := getSumCnt(t)

        return sum + (x - cnt) * t
    }

    return (f(right) - f(left - 1)) % MOD
}