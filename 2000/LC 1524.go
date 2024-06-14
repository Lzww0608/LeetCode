const MOD int = 1e9 + 7
func numOfSubarrays(arr []int) (ans int) {
    f := [2]int{1, 0}
    sum := 0
    for _, x := range arr {
        sum = sum ^ (x & 1)
        ans = (ans + f[sum ^ 1]) % MOD
        f[sum]++
    }

    return
}