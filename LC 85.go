func maximalRectangle(matrix [][]byte) int {
    n, m := len(matrix), len(matrix[0])
    dp := make([]int, m)
    ans := 0
    for i := 0; i < n; i++{
        for j := 0; j < m; j++{
            if matrix[i][j] == '1' {
                dp[j] += 1
            } else {
                dp[j] = 0
            }
        }
        ans = max(ans, largestRectangle(dp))
    }
    return ans
}

func largestRectangle(height []int) int {
    res, st := 0, []int{-1}
    if height[len(height)-1] != 0 {
        height = append(height, 0)
    }
    for i := 0; i < len(height); i++ {
        for len(st) > 1 && height[i] < height[st[len(st)-1]] {
            h := height[st[len(st)-1]]
            st = st[:len(st)-1]
            w := i - st[len(st)-1] - 1
            res = max(res, w * h)
        }
        st = append(st, i)
    }
    return res
}