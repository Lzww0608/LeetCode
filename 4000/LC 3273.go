func minDamage(power int, damage []int, health []int) int64 {
    ans := 0
    n := len(damage)
    id := make([]int, n)
    sum := 0
    for i := range id {
        id[i] = i
    }

    sort.Slice(id, func(i, j int) bool {
        x, y := (health[id[i]] + power - 1) / power, (health[id[j]] + power - 1) / power
        return float64(damage[id[i]]) / float64(x) < float64(damage[id[j]]) / float64(y)
    })

    for _, i := range id {
        sum += damage[i]
        time := (health[i] + power - 1) / power
        ans += sum * time
    }

    return int64(ans)
}