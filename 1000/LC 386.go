func lexicalOrder(n int) []int {
    ans := make([]int, 0, n)
    for i, j := 0, 1; i < n; i++ {
        ans = append(ans, j)
        if j * 10 <= n {
            j *= 10
        } else {
            for j + 1 > n || j % 10 == 9 {
                j /= 10
            }
            j++
        }
    }    

    return ans
}