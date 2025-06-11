func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    return calc(firstWord) + calc(secondWord) == calc(targetWord)
}

func calc(s string) (ans int) {
    for _, c := range s {
        ans = ans * 10 + int(c - 'a')
    }
    return
}