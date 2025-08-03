func countElements(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    ans := n 

    for i := 0; i < n && nums[i] == nums[0]; i++ {
        ans--
    }

    for i := n - 1; ans > 0 && nums[i] == nums[n - 1]; i-- {
        ans--
    }

    return ans
}