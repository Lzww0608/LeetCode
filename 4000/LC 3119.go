func maxPotholes(road string, budget int) (ans int) {
    n := len(road)
    a := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if road[i] == '.' {
            continue
        }
        j := i + 1
        for j < n && road[j] == 'x' {
            j++
        }
        a = append(a, j - i)
        i = j
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    for i := 0; i < len(a) && budget > 1; i++ {
        if budget >= a[i] + 1 {
            ans += a[i]
            budget -= a[i] + 1
        } else {
            ans += budget - 1
            break
        }
    }

    return
}