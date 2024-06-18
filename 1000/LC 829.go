func consecutiveNumbersSum(n int) (ans int) {
    for k := 1; k * k < n * 2; k++ {
        if (n * 2) % k == 0 && (2 * n / k - k + 1) & 1 == 0 {
            ans++
        }
    }

    return ans
}