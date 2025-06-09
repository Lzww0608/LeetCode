const N int = 26
func countQuadruples(a string, b string) (ans int) {
    pos := [N]int{}
    for i, c := range a {
        x := int(c - 'a')
        if pos[x] == 0 {
            pos[x] = i + 1
        }
    }

    mn := math.MaxInt32 
    for i, c := range b {
        x := int(c - 'a')
        if pos[x] > 0 {
            d := pos[x] - i
            if d < mn {
                ans = 1
                mn = d
            } else if mn == d {
                ans++
            }
        }
    }

    return
}