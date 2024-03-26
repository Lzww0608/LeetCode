func lengthOfLastWord(s string) int {
    ans, n := 0, len(s)
    i := n - 1
    for i >= 0 && s[i] == ' ' {
        i--
    }

    for i >= 0 && s[i] != ' ' {
        i--
        ans++
    }
    
    return ans
}