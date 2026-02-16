func countSubarrays(nums []int, k int64) int64 {
    mx := []int{}
    mn := []int{}

    ans := 0
    j := 0
    for i, x := range nums {
        for len(mx) > 0 && nums[mx[len(mx) - 1]] <= x {
            mx = mx[:len(mx) - 1]
        }
        mx = append(mx, i)

        for len(mn) > 0 && nums[mn[len(mn) - 1]] >= x {
            mn = mn[:len(mn) - 1]
        }
        mn = append(mn, i)

        for len(mx) > 0 && len(mn) > 0 && (nums[mx[0]] - nums[mn[0]]) * (i - j + 1) > int(k) {
            if mx[0] == j {
                mx = mx[1:]
            }
            if mn[0] == j {
                mn = mn[1:]
            }

            j++
        }
        ans += i - j + 1
    }

    return int64(ans)
}