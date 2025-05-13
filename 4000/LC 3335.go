const N int = 26
const MOD int = 1_000_000_007
func lengthAfterTransformations(s string, t int) (ans int) {
    cnt := [N]int{}
    for i := range s {
        x := int(s[i] - 'a')
        cnt[x]++
    }

    for i := 0; i < t; i++ {
        tmp := [N]int{}
        for i := 0; i < N - 1; i++ {
            tmp[i+1] += cnt[i]
        }
        tmp[0] += cnt[N-1]
        tmp[1] += cnt[N-1]
        for i := range cnt {
            cnt[i] = tmp[i] % MOD
        }
    }

    for _, x := range cnt {
        ans = (ans + x) % MOD
    }

    return
}