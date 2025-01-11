const MOD int = 1_000_000_007
func breakfastNumber(staple []int, drinks []int, x int) (ans int) {
    sort.Ints(staple)
    sort.Ints(drinks)
    m, n := len(staple), len(drinks)
    i, j := 0, n - 1

    for i < m && j >= 0 {
        if staple[i] + drinks[j] <= x {
            ans += j + 1
            ans %= MOD
            i++
        } else {
            j--
        }
    }
    
    return
}