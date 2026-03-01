var p, q int

func makeParityAlternating(nums []int) []int {
    p = slices.Min(nums)
    q = slices.Max(nums)
    if len(nums) == 1 {
        return []int{0, 0}
    }
    a := solve(nums, true)
    b := solve(nums, false)
    if a[0] < b[0] {
        return a
    } else if a[0] > b[0] {
        return b
    }

    if a[1] < b[1] {
        return a
    }
    return b
}

func solve(a []int, t bool) []int {
    ans := make([]int, 2)
    mn, mx := math.MaxInt32, math.MinInt32

    for i, x := range a {
        if (x & 1 == i & 1) == t {
            mn = min(mn, x)
            mx = max(mx, x)
        } else {
            ans[0]++
            if x == p {
                x++
            } else if x == q {
                x--
            }
            mn = min(mn, x)
            mx = max(mx, x)
        }
    }

    ans[1] = mx - mn 
    return ans
}