func countOppositeParity(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    a, b := 0, 0
    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        if x & 1 == 0 {
            ans[i] = b
            a++
        } else {
            ans[i] = a
            b++
        }
    }

    return ans
}