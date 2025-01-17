const N int = 1_000_001
func matrixMedian(grid [][]int) int {
    l, r := 0, N
    m, n := len(grid), len(grid[0])
    k := m * n / 2 + 1

    for l + 1 < r {
        cnt := 0
        mid := l + ((r - l) >> 1)
        for _, v := range grid {
            cnt += sort.SearchInts(v, mid + 1)
        }
        if cnt >= k {
            r = mid
        } else {
            l = mid 
        }
    }

    return r 
}