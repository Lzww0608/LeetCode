func stoneGameVIII(stones []int) int {
    n := len(stones)
    sum := 0
    for _, x := range stones {
        sum += x 
    }

    ans := sum 
    for i := n - 1; i > 1; i-- {
        sum -= stones[i]
        ans = max(ans, sum - ans)
    }

    return ans
}