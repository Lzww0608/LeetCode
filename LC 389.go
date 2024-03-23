func findTheDifference(s string, t string) byte {
    x := byte(t[len(t)-1])

    for i := 0; i < len(s); i++ {
        x ^= byte(s[i]) ^ byte(t[i])
    }

    return x

}