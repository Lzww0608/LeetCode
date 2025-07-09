func shuffle(nums []int, n int) []int {
    ans := make([]int, n * 2)
    for i := 0; i < n * 2; i++ {
        if i & 1 == 0 {
            ans[i] = nums[i / 2]
        } else {
            ans[i] = nums[i / 2 + n]
        }
    }

    return ans
}