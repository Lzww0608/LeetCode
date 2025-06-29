func meetRequirement(n int, lights [][]int, requirement []int) (ans int) {
    dif := make([]int, n + 1)
    for _, v := range lights {
        dif[max(v[0] - v[1], 0)]++
        dif[min(v[0] + v[1] + 1, n)]--
    }

    sum := 0
    for i := 0; i < n; i++ {
        sum += dif[i]
        if sum >= requirement[i] {
            ans++
        }
    }

    return
}