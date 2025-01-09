const MOD int = 1_000_000_007
func beautifulBouquet(flowers []int, cnt int) (ans int) {
    m := make(map[int]int)
    n := len(flowers)

    for l, r := 0, 0; r < n; r++ {
        x := flowers[r]
        m[x]++
        for m[x] > cnt {
            y := flowers[l]
            m[y]--
            l++
        }
        ans += r - l + 1
        ans %= MOD
    }

    return
}