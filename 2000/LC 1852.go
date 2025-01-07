func distinctNumbers(nums []int, k int) []int {
    m := make(map[int]int)
    n := len(nums)
    ans := make([]int, n - k + 1)
    for l, r := 0, 0; r < n; r++ {
        m[nums[r]]++
        if r - l + 1 == k {
            ans[l] = len(m)
            if m[nums[l]]--; m[nums[l]] == 0 {
                delete(m, nums[l])
            }
            l++
        }
    }

    return ans
}