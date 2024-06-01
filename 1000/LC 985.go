func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    n := len(nums)
    ans := make([]int, n)
    sum := 0
    for _, x := range nums {
        if x & 1 == 0 {
            sum += x
        }
    }
    for i, q := range queries {
        val, idx := q[0], q[1]
        if nums[idx] & 1 == 0 {
            sum -= nums[idx]
        }
        nums[idx] += val
        if nums[idx] & 1 == 0 {
            sum += nums[idx]
        }
        ans[i] = sum
    }

    return ans
    
}