func canBeTypedWords(text string, brokenLetters string) (ans int) {
    m := make(map[byte]bool)
    for i := range brokenLetters {
        m[brokenLetters[i]] = true
    }

    str := strings.Split(text, " ")
    next:
    for _, s := range str {
        for i := range s {
            if m[s[i]] {
                continue next
            }
        }
        ans++
    }

    return 
}