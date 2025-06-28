const N int = 10
func countPoints(rings string) (ans int) {
    cnt := [N]map[byte]bool{}
    for i := range cnt {
        cnt[i] = make(map[byte]bool)
    }

    n := len(rings)
    for i := 0; i < n; i += 2 {
        x := int(rings[i+1] - '0')
        cnt[x][rings[i]] = true
    }

    for i := range cnt {
        if len(cnt[i]) == 3 {
            ans++
        }
    }

    return 
}