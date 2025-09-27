const N int = 26
func rearrangeCharacters(s string, target string) (ans int) {
    cntS := [N]int{}
    cntT := [N]int{}
    for _, c := range s {
        cntS[int(c - 'a')]++
    }
    for _, c := range target {
        cntT[int(c - 'a')]++
    }

    ans = len(s)
    for i := range N {
        if cntT[i] == 0 {
            continue
        }
        ans = min(ans, cntS[i] / cntT[i])
    }

    return
}