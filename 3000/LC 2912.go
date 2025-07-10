const MOD int = 1_000_000_007
func numberOfWays(n int, m int, k int, source []int, dest []int) int {
    f := [4]int{}
    /*
    * 0: x1 != x2, y1 != y2
    * 1: x1 == x2, y1 != y2
    * 2: x1 != x2, y1 == y2
    * 3: x1 == x2, y1 == y2
    */
    if source[0] == dest[0] {
        if source[1] == dest[1] {
            f[3] = 1
        } else {
            f[1] = 1
        }
    } else {
        if source[1] == dest[1] {
            f[2] = 1
        } else {
            f[0] = 1
        }
    }

    for i := 1; i <= k; i++ {
        tmp := [4]int{f[0], f[1], f[2], f[3]}
        f[0] = (tmp[0] * (m + n - 4) + tmp[1] * (n - 1) + tmp[2] * (m - 1)) % MOD
        f[1] = (tmp[0] + tmp[1] * (m - 2) + tmp[3] * (m - 1)) % MOD
        f[2] = (tmp[0] + tmp[2] * (n - 2) + tmp[3] * (n - 1)) % MOD
        f[3] = (tmp[1] + tmp[2]) % MOD
    }

    return f[3]
}