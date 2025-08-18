func displayTable(orders [][]string) [][]string {
    m := make(map[string]map[string]int)
    foods := make(map[string]bool)
    for _, order := range orders {
        table, food := order[1], order[2]
        if _, ok := m[table]; !ok {
            m[table] = make(map[string]int)
        }
        m[table][food]++
        foods[food] = true
    }

    a := make([]string, 0, len(foods))
    for k := range foods {
        a = append(a, k)
    }
    sort.Strings(a)
    b := make([]int, 0, len(m))
    for k := range m {
        x, _ := strconv.Atoi(k)
        b = append(b, x)
    }
    sort.Ints(b)

    ans := make([][]string, len(m) + 1)
    for i := 0; i < len(ans); i++ {
        ans[i] = make([]string, len(a) + 1)
    }
    ans[0][0] = "Table"
    for i := 1; i <= len(a); i++ {
        ans[0][i] = a[i - 1]
    }

    for i := 1; i < len(ans); i++ {
        table := strconv.Itoa(b[0])
        ans[i][0] = table
        for j := 1; j <= len(a); j++ {
            ans[i][j] = strconv.Itoa(m[table][a[j - 1]])
        }
        b = b[1:]
    }

    return ans
}