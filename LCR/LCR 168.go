func nthUglyNumber(n int) int {
    g := []int{2, 3, 5}
    p := [3]int{}
    ugly := make([]int, 0, n)
    ugly = append(ugly, 1)

    for i := 1; i < n; i++ {
        mn := math.MaxInt32
        for j, x := range g {
            mn = min(mn, x * ugly[p[j]])
        }

        for j, x := range g {
            if mn == x * ugly[p[j]] {
                p[j]++
            }
        }

        ugly = append(ugly, mn)
    }

    return ugly[n - 1]
}