func minChanges(nums []int, k int) int {
    dif := make([]int, k + 2)
    n := len(nums)

    for i := 0; i < n / 2; i++ {
        x, y := nums[i], nums[n - i - 1]
        if x > y {
            x, y = y, x
        }

        d := y - x
        dif[0]++
        dif[d]--
        dif[d + 1]++
        mx := max(y, k - x)
        dif[mx + 1]++
    }

    ans := dif[0]
    for i := 1; i <= k + 1; i++ {
        dif[i] += dif[i - 1]
        ans = min(ans, dif[i])
    }

    return ans
}