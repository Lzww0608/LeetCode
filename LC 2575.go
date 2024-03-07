func divisibilityArray(word string, m int) []int {
    n := len(word)
    ans := make([]int, n)

    remainder := 0
    for i, c := range word {
        x := int(c - '0')
        if (remainder + x) % m == 0 {
            remainder = 0
            ans[i] = 1
        } else {
            remainder = (remainder + x) % m * 10
        }
    }
    return ans
}