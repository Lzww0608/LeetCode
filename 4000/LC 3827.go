func countMonobit(n int) int {
    ans := 0
    for i := 0; (1 << i) - 1 <= n; i++ {
        ans++
    }

    return ans
}