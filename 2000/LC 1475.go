func finalPrices(prices []int) []int {
    n := len(prices)
    ans := make([]int, n)
    s := []int{}
    
    for i := n - 1; i >= 0; i-- {
        for len(s) > 0 && prices[i] < s[len(s)-1] {
            s = s[:len(s)-1]
        }

        if len(s) > 0 {
            ans[i] = prices[i] - s[len(s)-1]
        } else {
            ans[i] = prices[i]
        }

        s = append(s, prices[i])
    }
    return ans
}
