func findErrorNums(nums []int) []int {
    n := len(nums)
    m := make([]int, n + 1)
    for _, x := range nums {
        m[x]++
    }

    ans := make([]int, 2)
    for i := 1; i <= n; i++ {
        if m[i] == 0 {
            ans[1] = i
        } else if m[i] == 2 {
            ans[0] = i
        }
    }
    return ans
}