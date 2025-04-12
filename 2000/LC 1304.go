func sumZero(n int) []int {
    ans := make([]int, n)
    sum := 0
    for i := 0; i < n - 1; i++ {
        ans[i] = i + 1
        sum += i + 1
    }
    ans[n-1] = -sum
    return ans
}