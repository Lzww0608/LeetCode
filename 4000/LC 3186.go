func maximumTotalDamage(power []int) int64 {
    sort.Ints(power)
    a := []int{}
    m := map[int]int{}
    for i, x := range power {
        if i == 0 || x != power[i-1] {
            a = append(a, x)
            m[len(a)-1] = x 
        } else {
            a[len(a)-1] += x 
        }
    }

    n := len(a)
    f := make([]int, n)
    for i := 0; i < n; i++ {
        x, j := m[i], i - 1
        f[i] = a[i]
        for j = i - 1; j >= max(0, i - 2) && m[j] >= x - 2; j-- {
        }
        if j >= 0 {
            f[i] = max(f[i], a[i] + f[j])
        }
        f[i] = max(f[i], f[max(0, i - 1)], f[max(0, i - 2)])
    }

    return int64(f[n-1])
}