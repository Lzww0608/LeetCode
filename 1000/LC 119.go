func getRow(rowIndex int) []int {
    ans := make([]int, rowIndex + 1)
    ans[0] = 1

    //C(n, m) = C(n, m - 1) * (n - m + 1) * m
    for i := 1; i <= rowIndex; i++ {
        ans[i] = ans[i-1] * (rowIndex - i + 1) / i
    }
    
    return ans
}
