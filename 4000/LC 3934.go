const MOD1 int = 1_070_777_777
const MOD2 int = 1_000_000_007

func smallestUniqueSubarray(nums []int) int {
    n := len(nums)
    

    base1 := 9e8 - rand.Intn(1e8)
    base2 := 9e8 - rand.Intn(1e8)
    type pair struct {h1, h2 int}
    powBase := make([]pair, n + 1)
    preHash := make([]pair, n + 1)
    powBase[0] = pair{1, 1}

    for i, c := range nums {
        powBase[i+1] = pair{powBase[i].h1 * base1 % MOD1, powBase[i].h2 * base2 % MOD2};
        preHash[i+1] = pair{(preHash[i].h1 * base1 + int(c)) % MOD1, (preHash[i].h2 * base2 + int(c)) % MOD2}
    }

    // target[l:r)
    subHash := func(l, r int) pair {
        h1 := ((preHash[r].h1 - preHash[l].h1 * powBase[r-l].h1) % MOD1 + MOD1) % MOD1
        h2 := ((preHash[r].h2 - preHash[l].h2 * powBase[r-l].h2) % MOD2 + MOD2) % MOD2
        return pair{h1, h2}
    }
    
    m := make(map[pair]int)
    check := func(mid int) bool {
        clear(m)
        for i := range n - mid + 1 {
            cur := subHash(i, i + mid)
            m[cur]++
        }

        for _, v := range m {
            if v == 1 {
                return true
            }
        }

        return false
    }

    l, r := 0, n
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}