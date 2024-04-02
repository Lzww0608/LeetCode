func shortestToChar(s string, c byte) []int {
    n := len(s)
    ans := make([]int, n)
    a := []int{}
    for i := range s {
        if s[i] == c {
            a = append(a, i)
        }
    }
    m := len(a)
    for i, j := 0, 0; i < n; i++ {
        if i > a[j] && j + 1 < m && i - a[j] >= a[j+1] - i {
            j++
        } 
        ans[i] = abs(i - a[j])
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}