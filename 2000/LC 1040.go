func numMovesStonesII(stones []int) []int {
    n := len(stones)
    sort.Ints(stones)
    mx := max(stones[n - 1] - stones[1], stones[n - 2] - stones[0]) - n + 2
    if stones[n - 1] - stones[1] == n - 2 || stones[n - 2] - stones[0] == n - 2 {
        return []int{min(2, mx), mx}
    }
    mn := mx 
    for l, r := 0, 0; r < n; r++ {
        for stones[r] - stones[l] >= n {
            l++
        }
        mn = min(mn, n - (r - l + 1))
    }

    return []int{mn, mx}
}