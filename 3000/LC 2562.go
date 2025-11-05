func findTheArrayConcVal(nums []int) int64 {
    ans := 0
    n := len(nums)

    for i, j := 0, n - 1; i <= j; i, j = i + 1, j - 1 {
        if i == j {
            ans += nums[i]
        } else {
            s := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
            x, _ := strconv.Atoi(s)
            ans += x
        }
    }

    return int64(ans)
}