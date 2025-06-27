func maxSubsequence(nums []int, k int) []int {
    n := len(nums)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return nums[id[i]] > nums[id[j]]
    })
    sort.Ints(id[:k])
    ans := make([]int, k)
    for i := range k {
        ans[i] = nums[id[i]]
    }

    return ans
}