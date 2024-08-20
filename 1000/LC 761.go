func makeLargestSpecial(s string) string {
    n := len(s)
    if n <= 2 {
        return s
    }

    subs := sort.StringSlice{}
    cnt, left := 0, 0
    for i, c := range s {
        if c == '1' {
            cnt++
        } else if cnt--; cnt == 0 {
            subs = append(subs, "1" + makeLargestSpecial(s[left+1:i]) + "0")
            left = i + 1
        }
    }
    sort.Sort(sort.Reverse(subs))
    return strings.Join(subs, "")
}