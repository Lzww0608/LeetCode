func removeVowels(s string) string {
    t := []byte{}
    for i := range s {
        if s[i] != 'a' && s[i] != 'e' && s[i] != 'o' && s[i] != 'u' && s[i] != 'i' {
            t = append(t, s[i])
        }
    } 

    return string(t)
}