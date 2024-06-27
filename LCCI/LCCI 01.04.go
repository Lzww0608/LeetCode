func canPermutePalindrome(s string) bool {
    m := map[rune]int{}
    for _, c := range s {
        m[c]++
    }

    odd := 0
    for _, x := range m {
        if x & 1 == 1 {
            odd++
            if odd > 1 {
                return false
            }
        }
    }

    return true
}