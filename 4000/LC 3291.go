const MOD int = 1_070_777_777
func minValidStrings(words []string, target string) (ans int) {
    n := len(target)
    base := 9e8 - rand.Intn(1e8)
    
    preBase := make([]int, n + 1)
    preHash := make([]int, n + 1)
    preBase[0] = 1
    for i, c := range target {
        preBase[i+1] = preBase[i] * base % MOD
        preHash[i+1] = (preHash[i] * base + int(c)) % MOD
    }

    subHash := func(l, r int) int {
        return ((preHash[r] - preHash[l] * preBase[r-l]) % MOD + MOD) % MOD
    }

    maxLen := 0
    for _, s := range words {
        maxLen = max(maxLen, len(s))
    }

    sets := make([]map[int]bool, maxLen)
    for i := range sets {
        sets[i] = make(map[int]bool)
    }

    for _, s := range words {
        h := 0
        for i, c := range s {
            h = (h * base + int(c)) % MOD
            sets[i][h] = true
        }
    }

    cur, nxt := 0, 0
    for i := 0; i < n; i++ {
        maxJump := sort.Search(min(n - i, maxLen), func(j int) bool {
            return !sets[j][subHash(i, i + j + 1)]
        })
        nxt = max(nxt, i + maxJump)
        if i == cur {
            if i == nxt {
                return -1
            }
            cur = nxt 
            ans++
        }
    }

    return
}