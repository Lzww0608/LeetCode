func decode(encoded []int) []int {
    xor_sum, n := 0, len(encoded) + 1
    for i := 1; i <= n; i++ {
        xor_sum ^= i
    }
    for i := 1; i < n-1; i += 2 {
        xor_sum ^= encoded[i]
    }
    ans := make([]int, n)
    ans[0] = xor_sum
    for i := 1; i < n; i++ {
        ans[i] = ans[i-1] ^ encoded[i-1]
    }
    return ans
}