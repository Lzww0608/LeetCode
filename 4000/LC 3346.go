func maxFrequency(nums []int, k int, numOperations int) (ans int) {
    sort.Ints(nums)
    n := len(nums)

    cnt := 0
    for l, r, i := 0, 0, 0; i < n; i++ {
        x := nums[i]
        cnt++
        if i < n - 1 && nums[i] == nums[i+1] {
            continue
        }

        for nums[l] < x - k {
            l++
        }

        for r < n && nums[r] <= x + k {
            r++
        }

        ans = max(ans, min(r - l, cnt + numOperations))
        cnt = 0
    }

    if ans >= numOperations {
        return
    }

    for l, r := 0, 0; r < n; r++ {
        for nums[l] < nums[r] - k * 2 {
            l++
        }
        ans = max(ans, min(numOperations, r - l + 1))
    }

    return
}