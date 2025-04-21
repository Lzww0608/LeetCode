func repeatedCharacter(s string) byte {
    vis := make([]bool, 26)
    for i := range s {
        x := int(s[i] - 'a')
        if vis[x] {
            return s[i]
        }
        vis[x] = true
    }

    return ' '
}