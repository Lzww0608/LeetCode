func findBlackPixel(picture [][]byte, target int) (ans int) {
    m, n := len(picture), len(picture[0])
    rowCnt := make([][]int, m)
    colCnt := make([][]int, n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if picture[i][j] == 'B' {
                rowCnt[i] = append(rowCnt[i], j)
                colCnt[j] = append(colCnt[j], i)
            }
        }
    }

next:
    for j := 0; j < n; j++ {
        if len(colCnt[j]) != target {
            continue
        }
        for _, i := range colCnt[j] {
            if len(rowCnt[i]) != target {
                continue next
            }
            for _, k := range rowCnt[i] {
                if picture[colCnt[j][0]][k] != picture[i][k] {
                    continue next
                }
            }
        }
        ans += target
    }

    return
}