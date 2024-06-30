func largeGroupPositions(s string) [][]int {
    ans := [][]int{}
    for i := 0; i + 1 < len(s); i++ {
        j := i
        for j + 1 < len(s) && s[j] == s[j+1] {
            j++
        }
        if j - i >= 2 {
            ans = append(ans, []int{i, j})
        }
        i = j
    }
    return ans
}
