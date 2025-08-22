func platesBetweenCandles(s string, queries [][]int) []int {
    n := len(s)
    ans := make([]int, len(queries))
    sum := make([]int, n + 1)
    left := make([]int, n)
    right := make([]int, n)
    p := -1
    for i := range s {
        sum[i + 1] = sum[i]
        if s[i] == '*' {
            sum[i + 1]++
        } else {
            p = i
        }
        left[i] = p
    }

    p = n - 1
    for i := n - 1; i >= 0; i-- {
        if s[i] == '|' {
            p = i
        } 
        right[i] = p
    }

    for i, q := range queries {
        l, r := right[q[0]], left[q[1]]
        if l < r {
            ans[i] = sum[r] - sum[l]
        }
    }

    return ans
}