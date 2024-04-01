func numberOfLines(widths []int, s string) []int {
    n, cnt := 1, 0
    for i := range s {
        t := widths[int(s[i] - 'a')]
        if cnt + t > 100 {
            n++
            cnt = t
        } else {
            cnt += t
        }
    } 
    return []int{n, cnt}
}