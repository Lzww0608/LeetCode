func transformArray(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    j := n - 1
    for _, x := range nums {
        if x & 1 == 1 {
            ans[j] = 1
            j--
        }
    }

    return ans
}