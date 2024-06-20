func numberOfArithmeticSlices(nums []int) (ans int) {
    n := len(nums)
    a := make([]map[int]int, n)
    for i := range nums {
        a[i] = make(map[int]int)
        for j := 0; j < i; j++ {
            d := nums[i] - nums[j]
            ans += a[j][d]
            a[i][d] += a[j][d] + 1
        }
    }

    return
}