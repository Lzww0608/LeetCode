func totalScore(hp int, damage []int, requirement []int) int64 {
    n := len(damage)
    ans := 0

    a := make([]int, 0, n)
    sum := 0
    for i := range n {
        a = append(a, hp)
        sum += damage[i]

        j := sort.SearchInts(a, sum + requirement[i])
        ans += len(a) - j

        hp += damage[i]
    }

    return int64(ans)
}