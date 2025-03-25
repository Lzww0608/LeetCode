func minimumSum(n int, k int) int {
    a := min(n, k / 2)
    n -= a 

    return a * (1 + a) / 2 + (k + k + n - 1) * n / 2 

}