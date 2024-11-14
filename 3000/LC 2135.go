func wordCount(startWords []string, targetWords []string) (ans int) {
    m := map[int]struct{}{}

    for _, s := range startWords {
        x := 0
        for i := range s {
            x |= 1 << (int(s[i] & 31) - 1)
        }
        m[x] = struct{}{}
    }

    for _, s := range targetWords {
        x := 0
        for i := range s {
            x |= 1 << (int(s[i] & 31) - 1)
        }

        for i := range s {
            t := x ^ (1 << (int(s[i] & 31) - 1))
            if _, ok := m[t]; ok {
                ans++
                break
            }
        }
    }

    return
}