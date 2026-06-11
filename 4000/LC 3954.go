func sumOfGoodIntegers(n int, k int) int {
    ans := 0
    for x := max(0, n - k); x <= n + k; x++ {
        if n & x == 0 {
            ans += x
        }
    }

    return ans
}