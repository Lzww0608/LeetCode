func minimumSeconds(nums []int) int {
    m := map[int][]int{}
    n := len(nums)
    for i, x := range nums {
        m[x] = append(m[x], i)
    }
    ans := n
    for _, v := range m {
        t := n - v[len(v)-1] + v[0]
        for i := 1; i < len(v); i++ {
            t = max(t, v[i] - v[i-1])
        }
        ans = min(ans, t)
    }
    return ans / 2
}
