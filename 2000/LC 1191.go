const MOD int = 1_000_000_007
func kConcatenationMaxSum(arr []int, k int) int {
    if k == 1 {
        sum, cur := 0, 0
        for _, x := range arr {
            cur = max(cur + x, x)
            sum = max(sum, cur)
        }
        return sum % MOD
    }

    sum := 0
    for _, x := range arr {
        sum += x
    }

    t := kConcatenationMaxSum(append(arr, arr...), 1)
    return max(sum * k, t, t + sum * (k - 2)) % MOD
}