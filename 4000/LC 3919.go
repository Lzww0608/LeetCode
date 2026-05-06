func minCost(nums []int, queries [][]int) []int {
    n := len(nums)

    pre := make([]int, n)
    suf := make([]int, n)
    for i := 1; i < n; i++ {
        if i < n - 1 {
            x := nums[i] - nums[i - 1]
            y := nums[i + 1] - nums[i]
            if x <= y {
                pre[i] = pre[i - 1] + 1
            } else {
                pre[i] = pre[i - 1] + x
            }
        } else {
            pre[i] = pre[i - 1] + 1
        }

        if i > 1 {
            x := nums[i - 1] - nums[i - 2]
            y := nums[i] - nums[i - 1]
            if x <= y {
                suf[i] = suf[i - 1] + y
            } else {
                suf[i] = suf[i - 1] + 1
            }
        } else {
            suf[i] = suf[i - 1] + 1
        }
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        l, r := q[0], q[1]
        if l < r {
            ans[i] = suf[r] - suf[l] 
        } else {
            ans[i] = pre[l] - pre[r]
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