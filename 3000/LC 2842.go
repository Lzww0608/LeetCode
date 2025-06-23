const MOD int = 1_000_000_007

const N int = 26

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }
    
    return res
}

func factor(x int) int {
    res := 1
    for i := 2; i <= x; i++ {
        res = res * i % MOD
    }

    return res
}

func comb(a, b int) int {
    return factor(b) * quickPow(factor(a) * factor(b - a) % MOD, MOD - 2) % MOD
}

type pair struct {
    freq, cnt int
}

func countKSubsequencesWithMaxBeauty(s string, k int) int {
    alp := [N]int{}
    for _, c := range s {
        alp[c - 'a']++
    }
    cnt := make(map[int]int)
    for _, x := range alp {
        if x > 0 {
            cnt[x]++
        }
    }

    a := []pair{}
    for k, v := range cnt {
        a = append(a, pair{k, v})
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i].freq > a[j].freq
    })

    ans := 1
    for _, v := range a {
        if v.cnt >= k {
            return ans * quickPow(v.freq, k) % MOD * comb(k, v.cnt) % MOD
        }
        ans = ans * quickPow(v.freq, v.cnt) % MOD
        k -= v.cnt
    }

    return 0
}