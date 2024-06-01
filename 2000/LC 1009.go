func bitwiseComplement(n int) int {
    ans := 1
    for ans < n {
        ans = (ans << 1) + 1
    }

    return ans ^ n
}
