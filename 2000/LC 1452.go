func peopleIndexes(favoriteCompanies [][]string) (ans []int) {
    idx := map[string]int{}
    tot := 0
    n := len(favoriteCompanies)
    a := make([][]int, n)
    for i, s := range favoriteCompanies {
        for _, t := range s {
            if _, ok := idx[t]; !ok {
                idx[t] = tot
                tot++
            }
            a[i] = append(a[i], idx[t])
        }
        sort.Ints(a[i])
    }
    for i := range favoriteCompanies {
        f := true
        for j := 0; j < n; j++ {
            if len(a[j]) < len(a[i]) || i == j {
                continue
            }
            p, q := 0, 0
            for p < len(a[i]) && q < len(a[j]) {
                if a[i][p] == a[j][q] {
                    p++
                } else {
                    q++
                }
            }
            if p == len(a[i]) {
                f = false
                break
            }
        }
        if f {
            ans = append(ans, i)
        }
    }

    return
}