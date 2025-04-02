const N int = 10
func isTransformable(s string, t string) bool {
    pos := [N][]int{}
    for i := range s {
        x := int(s[i] - '0')
        pos[x] = append(pos[x], i)
    }

    for i := range t {
        x := int(t[i] - '0')
        if len(pos[x]) == 0 {
            return false
        }

        for j := 0; j < x; j++ {
            if len(pos[j]) > 0 && pos[j][0] < pos[x][0] {
                return false
            }
        }
        pos[x] = pos[x][1:]
    }

    return true
}