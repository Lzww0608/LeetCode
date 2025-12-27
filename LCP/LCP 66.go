const N int = 26
func minNumBooths(demand []string) int {
    cnt := [N]int{}
    for _, s := range demand {
        cur := [N]int{}
        for _, c := range s {
            x := int(c - 'a')
            cur[x]++
        }
        for i := range N {
            cnt[i] = max(cnt[i], cur[i])
        }
    }

    ans := 0
    for _, x := range cnt {
        ans += x
    }

    return ans
}