func findOriginalArray(changed []int) []int {
    n := len(changed)
    if n & 1 == 1 {
        return []int{}
    }
    sort.Ints(changed)
    ans := make([]int, n / 2)
    m := map[int]int{}
    i := 0
    for _, x := range changed {
        if x & 1 == 0 && m[x/2] > 0 {
            ans[i] = x / 2
            i++
            m[x/2]--
        } else if m[x * 2] > 0 {
            ans[i] = x * 2
            i++
            m[x*2]--
        } else {
            m[x]++
        }
    }
    if i < n / 2 {
        return []int{}
    }

    return ans
}