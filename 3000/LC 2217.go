func kthPalindrome(queries []int, intLength int) []int64 {
    n := len(queries)
    ans := make([]int64, n)
    base := int(math.Pow10((intLength - 1) / 2))
    for i, q := range queries {
        if q > base * 9 {
            ans[i] = -1
            continue
        }
        left := base + q - 1
        x := left 
        if intLength & 1 == 1 {
            x /= 10
        }
        for x > 0 {
            left = left * 10 + x % 10
            x /= 10
        }
        ans[i] = int64(left)
    }

    return ans
}