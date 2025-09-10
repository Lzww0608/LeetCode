
import "math"
func minSubarraySort(nums []int, k int) []int {
    n := len(nums)
    ans := make([]int, n - k + 1)
    for i := 0; i < len(ans); i++ {
        ans[i] = solve(nums[i: i + k])
    }

    return ans
}

func solve(a []int) int {
    n := len(a)
    l, r := -1, -1
    mn, mx := math.MaxInt32, math.MinInt32
    for i := 0; i < n; i++ {
        if a[i] < mx {
            r = i
        } else {
            mx = a[i]
        }

        if a[n - i - 1] > mn {
            l = n - i - 1
        } else {
            mn = a[n - i - 1]
        }
    }

    if r > l {
        return r - l + 1
    }

    return 0
}