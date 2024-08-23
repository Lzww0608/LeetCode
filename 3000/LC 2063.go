func countVowels(word string) int64 {
    ans := 0
    n := len(word)
    m := map[rune]bool {
        'a': true,
        'e': true,
        'i': true,
        'o': true,
        'u': true,
    }
    for i, c := range word {
        if m[c] {
            ans += (i + 1) * (n - i)
        }
    }

    return int64(ans)
}