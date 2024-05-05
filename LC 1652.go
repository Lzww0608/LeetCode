func decrypt(code []int, k int) []int {
    n := len(code)
    ans := make([]int, n)
    if k == 0 {
        return ans
    }

    for j := range code {
        x := 0
        if k > 0 {
            for i := 1; i <= k; i++ {
                x += code[(i + j + n) % n]
            }
        } else {
            for i := -1; i >= k; i-- {
                x += code[(i + j + n) % n]
            }
        }
        ans[j] = x
    }

    return ans

}