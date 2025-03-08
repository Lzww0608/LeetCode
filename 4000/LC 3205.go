func maxScore(nums []int) (ans int) {
    n := len(nums)
    mx := 0
    for i := n - 1; i > 0; i-- {
        mx = max(nums[i], mx)
        ans += mx 
    }

    return 
}