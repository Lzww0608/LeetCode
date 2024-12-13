func minMoves(nums []int, limit int) int {
    dif := make([]int, limit * 2 + 2)
    n := len(nums)
    for i := 0; i < n / 2; i++ {
        x := nums[i] + nums[n-1-i]
        l, r := min(nums[i], nums[n-1-i]), max(nums[i], nums[n-1-i]) + limit
        dif[1] += 2
        dif[l+1] -= 1
        dif[x] -= 1
        dif[x+1] += 1
        dif[r+1] += 1
    }

    ans := math.MaxInt32
    cur := 0
    for i := 1; i <= limit * 2; i++ {
        cur += dif[i]
        ans = min(ans, cur)
    }

    return ans
}