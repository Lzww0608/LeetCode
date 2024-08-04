func countNumbersWithUniqueDigits(n int) int {
    //0: 1
    //1: 9
    //2: 9 * 9
    if n == 0 {
        return 1
    }
    ans := 10
    cur := 9
    for i := 2; i <= n; i++ {
        cur *= 10 - i + 1
        ans += cur
    }

    return ans
}