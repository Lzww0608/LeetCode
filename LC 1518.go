func numWaterBottles(a int, b int) int {
    ans, c := 0, 0
    for a > 0 {
        ans += a
        a, c = (a + c) / b, (a + c) % b
    }
    return ans
}