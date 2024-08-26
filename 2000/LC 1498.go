const MOD int = 1_000_000_007
func numSubseq(nums []int, target int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    l, r := 0, n - 1
    for l <= r {
        sum := nums[l] + nums[r]
        if sum > target {
            r--
        } else {
            ans = (ans + quickPow(2, r-l)) % MOD
            l++
        }
    }
    
    return ans
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % MOD
        }
        a = (a * a) % MOD
        r >>= 1
    }

    return res
}