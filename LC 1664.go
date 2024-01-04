func waysToMakeFair(nums []int) int {
    ans, n := 0, len(nums)
    prefix_even, prefix_odd := make([]int, n+1), make([]int, n+1)
    for i, x := range nums {
        if i & 1 == 0 {
            prefix_even[i+1] = prefix_even[i] + x
            prefix_odd[i+1] = prefix_odd[i] 
        } else {
            prefix_even[i+1] = prefix_even[i]
            prefix_odd[i+1] = prefix_odd[i] + x
        }
    }
    for i := 0; i < n; i++ {
        if i & 1 == 0 {
            if (prefix_even[i] + prefix_odd[n] - prefix_odd[i+1]) * 2 ==  prefix_even[n] + prefix_odd[n] - nums[i] {
                ans++
            } 
        } else {
            if (prefix_odd[i] + prefix_even[n] - prefix_even[i+1]) * 2 ==  prefix_even[n] + prefix_odd[n] - nums[i] {
                ans++
            } 
        }
    }
    return ans
}