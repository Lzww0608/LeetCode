func supplyWagon(a []int) []int {
    n := len(a)
    for i := range (n + 1) / 2 {
        idx := 0
        for j := range n - i - 1 {
            if a[j] + a[j + 1] < a[idx] + a[idx + 1] {
                idx = j
            }
        }
        a[idx] += a[idx + 1]
        a = append(a[:idx + 1], a[idx + 2:]...)
    }

    return a
}