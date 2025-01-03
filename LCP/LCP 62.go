func transportationHub(path [][]int) int {
    in := make(map[int]int)
    out := make(map[int]int)
    id := make(map[int]int)
    k := 0

    for _, v := range path {
        a, b := v[0], v[1]
        if _, ok := id[a]; !ok {
            id[a] = k
            k++
        }
        if _, ok := id[b]; !ok {
            id[b] = k
            k++
        }

        u, v := id[a], id[b]
        in[v]++
        out[u]++
    }

    for t, i := range id {
        if in[i] == k - 1 && out[i] == 0 {
            return t
        }
    }

    return -1
}