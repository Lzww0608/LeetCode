const N int = 26
func rankTeams(votes []string) string {
    f := [N][N]int{}
    for _, vote := range votes {
        for i, c := range vote {
            x := int(c - 'A')
            f[x][i]++
        }
    }

    a := []byte(votes[0])
    sort.Slice(a, func(i, j int) bool {
        x, y := int(a[i] - 'A'), int(a[j] - 'A')
        for i := 0; i < N; i++ {
            if f[x][i] != f[y][i] {
                return f[x][i] > f[y][i]
            }
        }
        return x < y
    })

    return string(a)
}