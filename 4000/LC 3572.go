func maxSumDistinctTriplet(x []int, y []int) int {
    n := len(x)
    m := make(map[int]int)
    for i := range n {
        m[x[i]] = max(m[x[i]], y[i])
    }

    if len(m) < 3 {
        return -1
    }
    a, b, c := -1, -1, -1
    for _, t := range m {
        if t > a {
            a, b, c = t, a, b
        } else if t > b {
            b, c = t, b
        } else if t > c {
            c = t
        }
    }

    return a + b + c
}