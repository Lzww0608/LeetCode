func dismantlingAction(arr string) byte {
    m := [26]int{}
    pos := [26]int{}
    for i, c := range arr {
        x := int(c - 'a')
        m[x]++
        if pos[x] == 0 {
            pos[x] = i + 1
        }
    }
    
    var ans byte = ' '
    mn := len(arr) + 1
    for i, x := range pos {
        if x > 0 && m[i] == 1 && x < mn  {
            mn = x
            ans = byte('a' + i)
        }
    }

    return ans
}
