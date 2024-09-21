func findLonelyPixel(picture [][]byte) (ans int) {
    m, n := len(picture), len(picture[0])
    row := make([]int, m)

    for i := 0; i < m; i++ {
        cnt := 0
        for j := 0; j < n; j++ {
            if picture[i][j] == 'B' {
                cnt++
                if cnt > 1 {
                    break
                }
            }
        }
        row[i] = cnt
    }

    for j := 0; j < n; j++ {
        cnt, pos := 0, -1
        for i := 0; i < m; i++ {
            if picture[i][j] == 'B' {
                cnt++
                pos = i
                if cnt > 1 {
                    break
                }
            }
        }
        if cnt == 1 && row[pos] == 1 {
            ans++
        }
    }

    return
}