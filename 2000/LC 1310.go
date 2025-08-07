func xorQueries(arr []int, queries [][]int) []int {
    ans := make([]int, len(queries))
    a := make([]int, len(arr) + 1)
    for i, x := range arr {
        a[i + 1] = a[i] ^ x
    }

    for i, query := range queries {
        l, r := query[0], query[1]
        ans[i] = a[r + 1] ^ a[l]
    }

    return ans
}