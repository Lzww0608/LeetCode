
func continuousSubarrays(nums []int) int64 {
    ans := 0
    n := len(nums)

    cnt := map[int]int{}
    mx, mn := 0, 0

    for l, r := 0, 0; r < n; r++ {
        x := nums[r]
        cnt[x]++

        for {
            mx, mn = x, x
            for k := range cnt {
                mx = max(mx, k)
                mn = min(mn, k)
            }
            if mx - mn <= 2 {
                break
            }

            y := nums[l]
            if cnt[y]--; cnt[y] == 0 {
                delete(cnt, y)
            }
            l++
        }

        ans += r - l + 1
    }

    return int64(ans)
}