const MOD int = 1_000_000_007
func createSortedArray(instructions []int) (ans int) {
    mx := slices.Max(instructions)
    f := make([]int, mx + 1)

    update := func(x int) {
        for i := x; i <= mx; i += i & (-i) {
            f[i]++
        }
    }

    query := func(x int) (res int) {
        for i := x; i > 0; i -= i & (-i) {
            res += f[i];
        }

        return
    }

    for i, x := range instructions {
        l := query(x - 1)
        r := i - query(x)
        ans += min(l, r)
        update(x)
    }

    return ans % MOD
}