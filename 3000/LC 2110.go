func getDescentPeriods(prices []int) (ans int64) {
    var prev int64 = 1
    for i := 1; i < len(prices); i++ {
        var cur int64 = 1
        if prices[i] == prices[i-1] - 1 {
            cur += prev
        }
        ans += cur
        prev = cur
    }


    return ans + 1
}