func countNumbers(cnt int) []int {
    var k int = int(math.Pow(10, float64(cnt)))
    ans := make([]int, 0, k);
    for i := 1; i < k; i++ {
        ans = append(ans, i)
    }

    return ans
}
