func countCoveredBuildings(n int, buildings [][]int) (ans int) {
    cnt := make(map[int]int)
    row := make([][]int, n)
    col := make([][]int, n)
    for _, v := range buildings {
        i, j := v[0] - 1, v[1] - 1
        row[i] = append(row[i], j)
        col[j] = append(col[j], i)
    }

    for k := range n {
        sort.Ints(row[k])
        sort.Ints(col[k])

        for i := 1; i < len(row[k]) - 1; i++ {
            cnt[k * n + row[k][i]]++
        }
        for i := 1; i < len(col[k]) - 1; i++ {
            cnt[col[k][i] * n + k]++
        }
    }

    for _, x := range cnt {
        if x == 2 {
            ans++
        }
    }

    return
}