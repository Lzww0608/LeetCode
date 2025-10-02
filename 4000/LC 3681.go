const N int = 32
func maxXorSubsequences(nums []int) (ans int) {
    basis := make([]int, N)
    for _, x := range nums {
        for i := N - 1; i >= 0; i-- {
            if x & (1 << i) == 0 {
                continue
            }

            if basis[i] == 0 {
                basis[i] = x 
                break
            }

            x ^= basis[i]
        }
    }

    for i := N - 1; i >= 0; i-- {
        ans = max(ans, ans ^ basis[i])
    }

    return
}