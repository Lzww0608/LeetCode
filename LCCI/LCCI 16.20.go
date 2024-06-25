func getValidT9Words(num string, words []string) (ans []string) {
    m := map[rune]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    for _, s := range words {
        f := true
        for i, c := range num {
            if !strings.Contains(m[c], string(s[i])) {
                f = false
                break
            }
        }
        if f {
            ans = append(ans, s)
        }
    }

    return
}