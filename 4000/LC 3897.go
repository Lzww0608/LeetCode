const MOD int = 1_000_000_007
func maxValue(a []int, b []int) (ans int) {
    n := len(a)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        if b[id[i]] == 0 || b[id[j]] == 0 {
            return b[id[i]] == 0
        } else if a[id[i]] == a[id[j]] {
            return b[id[i]] < b[id[j]]
        }
        return a[id[i]] > a[id[j]] 
    })

    for _, i := range id {
        for range a[i] {
            ans = (ans * 2 + 1) % MOD;
        }
        for range b[i] {
            ans = (ans << 1) % MOD;
        }
    }

    return
}