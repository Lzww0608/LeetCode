const MOD int = 1_000_000_007
func specialTriplets(nums []int) (ans int) {
    suf := make(map[int]int)
    pre := make(map[int]int)
    for _, x := range nums {
        suf[x]++
    }

    for _, x := range nums {
        suf[x]--
        ans = (ans + pre[x * 2] * suf[x * 2]) % MOD
        pre[x]++
    }

    return 
}