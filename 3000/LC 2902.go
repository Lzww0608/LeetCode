const MOD int = 1_000_000_007
func countSubMultisets(nums []int, l int, r int) (ans int) {
    cnt := make(map[int]int)
    sum := 0
    for _, x := range nums {
        cnt[x]++
        sum += x
    }

    if l > sum {
        return 0
    }

    r = min(r, sum)
    f := make([]int, r + 1)
    f[0] = cnt[0] + 1
    delete(cnt, 0)

    sum = 0
    for k, v := range cnt {
        sum = min(sum + k * v, r)
        for j := k; j <= sum; j++ {
            f[j] = (f[j] + f[j - k]) % MOD
        }
        for j := sum; j >= k * (v + 1); j-- {
            f[j] = (f[j] - f[j - k * (v + 1)] + MOD) % MOD
        }
    }

    for _, x := range f[l:r+1] {
        ans = (ans + x) % MOD
    }

    return
}