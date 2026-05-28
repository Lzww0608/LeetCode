func minOperations(a []int) int {
    n := len(a)

    check := func(delta int) int {
        k := slices.Index(a, 0)
        if delta == -1 {
            k = slices.Index(a, n - 1)
        }
        for j := range n - 1 {
            i := j + k    
            if a[i % n] + delta != a[(i + 1) % n] {
                return -1
            }
        }

        return k
    }

    x, y := check(1), check(-1)
    if x != -1 {
        return min(x, 2 + n - x)
    } else if y != -1 {
        return min(y + 1, n - y + 1)
    }

    return -1
}
