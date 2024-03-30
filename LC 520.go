func detectCapitalUse(word string) bool {
    cnt, n := 0, len(word)
    for i := range word {
        if word[i] >= 'A' && word[i] <= 'Z' {
            cnt++
        }
    }
    return cnt == 0 || cnt == 1 && word[0] >= 'A' && word[0] <= 'Z' || cnt == n
}