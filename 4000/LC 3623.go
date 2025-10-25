const MOD int = 1_000_000_007
func countTrapezoids(points [][]int) int {
    m := make(map[int]int)
    for _, point := range points {
        m[point[1]]++
    }

    ans := 0
    a := make([]int, 0, len(m))
    sum := 0
    for _, v := range m {
        a = append(a, v * (v - 1) / 2)
        sum = (sum + v * (v - 1) / 2) % MOD
    }

    for i := range a {
        ans = (ans + a[i] * (sum - a[i] + MOD) % MOD) % MOD
    }

    return ans * quickPow(2, MOD - 2) % MOD
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }

    return res
}