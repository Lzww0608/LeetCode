func longestSubarray(nums []int, limit int) (ans int) {
    mn, mx := []int{}, []int{}
    n := len(nums)

    for l, r := 0, 0; r < n; r++ {
        x := nums[r]

        for len(mx) > 0 && nums[mx[len(mx)-1]] < x {
            mx = mx[:len(mx)-1]
        }
        for len(mn) > 0 && nums[mn[len(mn)-1]] > x {
            mn = mn[:len(mn)-1]
        }

        mx = append(mx, r)
        mn = append(mn, r)
        for len(mx) > 0 && len(mn) > 0 && nums[mx[0]] - nums[mn[0]] > limit {
            if mx[0] == l {
                mx = mx[1:]
            }
            if mn[0] == l {
                mn = mn[1:]
            }
            l++
        }

        ans = max(ans, r - l + 1)
    }

    return ans
}