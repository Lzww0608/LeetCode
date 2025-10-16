const N int = 26
func checkDistances(s string, distance []int) bool {
    pos := [N]int{}
    for i := range pos {
        pos[i] = -1
    }

    for i := range s {
        x := int(s[i] - 'a')
        if pos[x] == -1 {
            pos[x] = i
        } else {
            if i - pos[x] - 1 != distance[x] {
                return false
            }
        }
    }

    return true
}