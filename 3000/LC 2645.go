func addMinimum(word string) int {
    valid := []rune{'a', 'b', 'c'}
    ans, i := 0, 0
    for _, x := range word {
        for x != valid[i] {
            i = (i + 1) % 3
            ans++
        }
        i = (i + 1) % 3
    } 
    ans += int('c' - word[len(word)-1])
    return ans
}