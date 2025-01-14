func answerQueries(nums []int, queries []int) []int {
    n := len(queries)
    ans := make([]int, n)
    sort.Ints(nums)

    for i := 1; i < len(nums); i++ {
        nums[i] += nums[i-1]
    }
    for i, q := range queries {
        ans[i] = sort.SearchInts(nums, q + 1)
    }

    return ans
}