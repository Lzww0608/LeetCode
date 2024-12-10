
import "math"

const MOD int = 1_070_777_777
func minimumCost(target string, words []string, costs []int) int {
    n := len(target)

    base := 9e8 - rand.Intn(1e8)
    powBase := make([]int, n + 1)
    preHash := make([]int, n + 1)
    powBase[0] = 1

    for i, c := range target {
        powBase[i+1] = powBase[i] * base % MOD;
        preHash[i+1] = (preHash[i] * base + int(c)) % MOD
    }

    // target[l:r)
    subHash := func(l, r int) int {
        return ((preHash[r] - preHash[l] * powBase[r-l]) % MOD + MOD) % MOD
    }

    minCost := make([]map[int]int, n + 1)
    lens := map[int]struct{}{}
    for i, word := range words {
        m := len(word)
        lens[m] = struct{}{}
        // word hash
        h := 0
        for _, c := range word {
            h = (h * base + int(c)) % MOD
        }

        if minCost[m] == nil {
            minCost[m] = map[int]int{}
        }
        if minCost[m][h] == 0 {
            minCost[m][h] = costs[i]
        } else {
            minCost[m][h] = min(minCost[m][h], costs[i])
        }
    }

    sortedLens := make([]int, 0, len(lens))
    for l := range lens {
        sortedLens = append(sortedLens, l)
    }
    slices.Sort(sortedLens)

    f := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        f[i] = math.MaxInt32 
        for _, sz := range sortedLens {
            if sz > i {
                break
            }

            if cost, ok := minCost[sz][subHash(i - sz, i)]; ok {
                f[i] = min(f[i], f[i-sz] + cost)
            }
        }
    }
    if f[n] >= math.MaxInt32 {
        return -1
    }
    return f[n]
}