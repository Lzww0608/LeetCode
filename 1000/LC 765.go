func minSwapsCouples(row []int) (ans int) {
    n := len(row)
    pos := make([]int, n)
    for i, x := range row {
        pos[x] = i
    }

    for i := 0; i < n; i += 2 {
        x, y := row[i], row[i+1]
        if x ^ 1 == y {
            continue
        }
        z := x ^ 1
        row[pos[z]] = y 
        pos[y] = pos[z]
        ans++
    }

    return
}