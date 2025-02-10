func largestEvenSum(nums []int, k int) int64 {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    sum := 0
    for _, x := range nums[:k] {
        sum += x
    } 
    if sum & 1 == 0 {
        return int64(sum)
    } 
    even, odd := -1, -1
    for _, x := range nums[k:] {
        if even != -1 && odd != -1 {
            break
        } else if even == -1 && x & 1 == 0 {
            even = x
        } else if odd == -1 && x & 1 == 1 {
            odd = x
        }
    }

    ans := -1
    for i := k - 1; i >= 0; i-- {
        if nums[i] & 1 == 1 && even != -1 {
            ans = max(ans, sum - nums[i] + even)
        } else if nums[i] & 1 == 0 && odd != -1 {
            ans = max(ans, sum - nums[i] + odd)
        }
    }

    return int64(ans)
}