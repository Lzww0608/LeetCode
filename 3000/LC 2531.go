const N int = 26
func isItPossible(s string, t string) bool {
    cnt1 := make(map[int]int)
    cnt2 := make(map[int]int)
    for i := range s {
        x := int(s[i] & 31 - 1)
        cnt1[x]++
    }
    for i := range t {
        x := int(t[i] & 31 - 1)
        cnt2[x]++
    }

    for i, x := range cnt1 {
        for j, y := range cnt2 {
            if i == j {
                if len(cnt1) == len(cnt2) {
                    return true
                }
            } else if len(cnt1) - b2i(x == 1) + b2i(cnt1[j] == 0) == 
                    len(cnt2) - b2i(y == 1) + b2i(cnt2[i] == 0) {
                        return true
            }
        }
    }
    
    return false
}

func b2i(x bool) int {
    if x {
        return 1
    }
    return 0
}

