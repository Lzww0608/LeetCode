func maxSum(nums []int, threshold []int) int64 {
    ans := 0
    n := len(nums)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return threshold[id[i]] < threshold[id[j]]
    })

    for j, i := range id {
        if threshold[i] > j + 1 {
            break 
        }
        ans += nums[i]
    } 

    return int64(ans)
}