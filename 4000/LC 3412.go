const N int = 26
func calculateScore(s string) int64 {
    ans := 0
    
    pos := [N][]int{}
    for i, c := range s {
        x := int(c - 'a')
        y := N - x - 1
        if len(pos[x]) > 0 {
            ans += i - pos[x][len(pos[x]) - 1]
            pos[x] = pos[x][:len(pos[x]) - 1]
        } else {
            pos[y] = append(pos[y], i)
        }
    }

    return int64(ans)
}