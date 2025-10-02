func solveQueries(nums []int, queries []int) []int {
    m := len(queries)
    mx := slices.Max(nums)
    pos := make([][]int, mx + 1)
    for i, x := range nums {
        pos[x] = append(pos[x], i)
    }

    ans := make([]int, m)
    n := len(nums)
    for i, q := range queries {
        x := nums[q]
        if len(pos[x]) == 1 {
            ans[i] = -1
        } else {
            k := sort.SearchInts(pos[x], q)
            l, r := (k - 1 + len(pos[x])) % len(pos[x]), (k + 1) % len(pos[x])
            ans[i] = min((pos[x][r] - pos[x][k] + n) % n, (pos[x][k] - pos[x][l] + n) % n)
        }
    }

    return ans
}