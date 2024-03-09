func getHint(secret string, guess string) string {
    m := map[int]int{}
    a, b := 0, 0
    for i := range secret {
        if secret[i] == guess[i] {
            a++
        } else {
            m[int(secret[i]-'0')]++
        }
        
    }

    for i := range secret {
        if secret[i] != guess[i] && m[int(guess[i]-'0')] > 0 {
            b++
            m[int(guess[i]-'0')]--
        }
    }

    return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}