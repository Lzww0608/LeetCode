func maxStudents(seats [][]byte) int {
    n := len(seats)
    a := make([]int, n)
    for i, seat := range seats {
        for j := range seat {
            if seat[j] == '.' {
                a[i] |= 1 << j
            }
        }
    }

    memo := make(map[[3]int]int)
    var dfs func(int, int, int) int 
    dfs = func(i, j, k int) (res int) {
        t := [3]int{i, j, k}
        if v, ok := memo[t]; ok {
            return v
        }
        defer func() {memo[t] = res} ()

        if j == 0 {
            if i == 0 {
                return 0
            }
            return dfs(i - 1, a[i - 1] &^ ((k << 1) | (k >> 1)), 0)
        }

        lb := j & (-j)
        return max(dfs(i, j ^ lb, k), dfs(i, j &^ (lb *3), k | lb) + 1)
    }

    return dfs(n - 1, a[n - 1], 0)
}