func absDifference(nums []int, k int) (ans int) {
    n := len(nums)
    sort.Ints(nums)

    for i := range k {
        ans += nums[i] - nums[n - i - 1]
    }

    return abs(ans)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    
    return x
}