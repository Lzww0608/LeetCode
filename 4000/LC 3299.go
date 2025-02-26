const MOD int = 1_000_000_007
func getSum(nums []int) (ans int) {
    inc := make(map[int]int)
    a := make(map[int]int)
    des := make(map[int]int)
    b := make(map[int]int)

    sum := 0
    for _, x := range nums {
        inc[x] = (inc[x] + inc[x-1] + (a[x-1] + 1) * x) % MOD
        a[x] = (a[x] + a[x-1] + 1) % MOD
        des[x] = (des[x] + des[x+1] + (b[x+1] + 1) * x) % MOD
        b[x] = (b[x] + b[x+1] + 1) % MOD
        sum = (sum + x) % MOD
    }

    for _, x := range inc {
        ans = (ans + x) % MOD
    }
    for _, x := range des {
        ans = (ans + x) % MOD
    }

    return (ans - sum + MOD) % MOD
}