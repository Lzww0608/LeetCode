func findWords(words []string) (ans []string) {
    m := map[byte]int {
        'q': 0,
        'w': 0,
        'e': 0,
        'r': 0,
        't': 0,
        'y': 0,
        'u': 0,
        'i': 0,
        'o': 0,
        'p': 0,
        'a': 1,
        's': 1,
        'd': 1,
        'f': 1,
        'g': 1,
        'h': 1,
        'j': 1,
        'k': 1,
        'l': 1,
        'z': 2,
        'x': 2,
        'c': 2,
        'v': 2,
        'b': 2,
        'n': 2,
        'm': 2,
    }
    for _, word := range words {
        t := m[toLower(word[0])]
        f := true
        for i := range word {
            if t != m[toLower(word[i])] {
                f = false
                break
            }
        }
        if f {
            ans = append(ans, word)
        }
    }
    return 
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return byte(c + 32)
    }
    return c
}
