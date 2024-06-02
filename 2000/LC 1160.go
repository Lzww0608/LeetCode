func countCharacters(words []string, chars string) (ans int) {
    m := map[int]int{}

    for i := range chars {
        m[int(chars[i] - 'a')]++
    }

    for _, s := range words {
        cnt := make([]int, 26)
        f := true
        for i := range s {
            t := int(s[i] - 'a')
            cnt[t]++
            if cnt[t] > m[t] {
                f = false
                break
            }
        }
        if f {
            ans += len(s)
        }
    }
    
    return 
}
