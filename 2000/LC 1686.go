func stoneGameVI(a []int, b []int) int {
    n := len(a)
    v := make([]int, n)
    for i := range v {
        v[i] = i
    }
    sort.Slice(v, func(i, j int) bool {
        return a[v[i]] + b[v[i]] > a[v[j]] + b[v[j]]
    })

    sumA, sumB := 0, 0
    for i, x := range v {
        if i & 1 == 0 {
            sumA += a[x]
        } else {
            sumB += b[x]
        }
    }

    if sumA > sumB {
        return 1
    } else if sumA < sumB {
        return -1
    }

    return 0
}