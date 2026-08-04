func countValidPrefixes(s string) (ans int) {
    zero, one := 0, 0
    for i := range s {
        x := int(s[i] - '0')
        if x == 0 {
            zero++
        } else {
            one++
        }

        if abs(one - zero) <= 1 {
            ans++
        }
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}