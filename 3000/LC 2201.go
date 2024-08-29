func digArtifacts(n int, artifacts [][]int, dig [][]int) (ans int) {
    a := make([][]int, n)
    for i := range a {
        a[i] = make([]int, n)
    }

    for _, v := range dig {
        a[v[0]][v[1]] = 1
    }

    for _, v := range artifacts {
        x1, y1 := v[0], v[1]
        x2, y2 := v[2], v[3]
        f := true
        for i := x1; i <= x2; i++ {
            for j := y1; j <= y2; j++ {
                if a[i][j] == 0 {
                    f = false
                    break
                }
            }
        }
        if f {
            ans++
        }
    }

    return 
}