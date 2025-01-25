func decode(encoded []int, first int) []int {
    n := len(encoded)
    ans := make([]int, n + 1)
    ans[0] = first
    for i, x := range encoded {
        ans[i+1] = ans[i] ^ x
    }

    return ans
}