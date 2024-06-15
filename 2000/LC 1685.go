func getSumAbsoluteDifferences(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    for i := 1; i < n; i++ {
        x := nums[i] - nums[i-1]
        ans[i] = ans[i-1] + x * i
    }

    tmp := make([]int, n)
    for i := n - 2; i >= 0; i-- {
        x := nums[i+1] - nums[i]
        tmp[i] = tmp[i+1] + x * (n - i - 1)
        ans[i] += tmp[i]
    }

    return ans
}