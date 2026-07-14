const MOD int = 1_000_000_007
func minimumCost(nums []int, k int) (ans int) {
    sum := 0
    for _, x := range nums {
        sum += x
    }

    cnt := (sum - 1) / k % MOD

    return (cnt * (cnt + 1) / 2) % MOD
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD 
        }
        r >>= 1
        a = a * a % MOD 
    }

    return res
}
