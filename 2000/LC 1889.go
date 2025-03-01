const MOD int = 1_000_000_007
func minWastedSpace(packages []int, boxes [][]int) int {
    sort.Ints(packages)
    m := len(packages)

    ans := math.MaxInt
    for _, box := range boxes {
        sort.Ints(box)
        n := len(box)
        if box[n-1] < packages[m-1] {
            continue
        }
        cur, l := 0, 0
        for _, x := range box {
            pos := sort.SearchInts(packages[l:], x + 1) + l
            cur += (pos - l) * x 
            l = pos
            if l >= m {
                break
            }
        }
        ans = min(ans, cur)
    }

    if ans == math.MaxInt {
        return -1
    }
    for _, x := range packages {
        ans -= x
    }

    return ans % MOD
}