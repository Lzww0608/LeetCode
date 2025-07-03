func minimumSwaps(nums []int) int {
    n := len(nums)
    mn, mx := 0, 0 
    for i, x := range nums {
        if x >= nums[mx] {
            mx = i
        }

        if x < nums[mn] {
            mn = i
        }
    }

    ans := n - 1 - mx + mn 
    if mx < mn {
        ans -= 1
    }

    return ans
}