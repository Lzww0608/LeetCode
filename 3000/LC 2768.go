func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
    cnt := make(map[int]int)
    for _, coordinate := range coordinates {
        i, j := coordinate[0], coordinate[1]
        if i - 1 >= 0 {
            if j - 1 >= 0 {
                cnt[(i - 1) * n + (j - 1)]++
            }
            if j + 1 < n {
                cnt[(i - 1) * n + j]++
            }
        } 

        if i + 1 < m {
            if j - 1 >= 0 {
                cnt[i * n + (j - 1)]++
            }
            if j + 1 < n {
                cnt[i * n + j]++
            }
        }
    } 

    ans := make([]int64, 5)
    sum := 0
    for _, v := range cnt {
        ans[v]++
        sum += v
    }
    ans[0] = int64(m * n - m - n + 1 - len(cnt))
    return ans
}