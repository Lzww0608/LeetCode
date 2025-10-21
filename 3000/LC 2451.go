func oddString(words []string) string {
    n := len(words)
    if n <= 2 {
        return words[0]
    }

    m := len(words[0])
    for i := 0; i < n; i++ {
        j, k := (i + 1) % n, (i + 2) % n 
        a := make([]int, m - 1)
        b := make([]int, m - 1)
        c := make([]int, m - 1)
        for l := 1; l < m; l++ {
            a[l - 1] = int(words[i][l] - 'a') - int(words[i][l - 1] - 'a')
            b[l - 1] = int(words[j][l] - 'a') - int(words[j][l - 1] - 'a')
            c[l - 1] = int(words[k][l] - 'a') - int(words[k][l - 1] - 'a')
        }

        b1 := reflect.DeepEqual(a, b)
        b2 := reflect.DeepEqual(b, c)
        if b1 && b2 {
            continue
        } else if b1 {
            return words[(i + 2) % n]
        } else if b2 {
            return words[i]
        } else {
            return words[(i + 1) % n]
        }
    }

    return ""
}