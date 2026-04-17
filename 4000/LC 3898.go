func findDegrees(matrix [][]int) []int {
    n := len(matrix)
    ans := make([]int, n)
    for i := range n {
        for j := range n {
            ans[i] += matrix[i][j]
        }
    }

    return ans
}