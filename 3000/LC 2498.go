func maxJump(stones []int) (ans int) {
    n := len(stones)
    if n <= 2 {
        return stones[n-1] - stones[0]
    }
    for i := 2; i < n; i++ {
        ans = max(ans, stones[i] - stones[i-2])
    }

    return
}