const MOD int = 1_000_000_007
const N int = 100_005
func sumOfFlooredPairs(nums []int) (ans int) {
    cnt := make([]int, N)
    for _, x := range nums {
        cnt[x]++
    }

    pre := make([]int, N + 1)
    for i := 1; i < N; i++ {
        pre[i+1] = pre[i] + cnt[i]
    }

    for k, v := range cnt {
        if v == 0 {
            continue
        }
        for d := 1; d * k < N; d++ {
            ans = (ans + v * d * (pre[min(N - 1, (d + 1) * k)] - pre[d * k])) % MOD
        }
    }

    return
}