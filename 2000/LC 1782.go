func countPairs(n int, edges [][]int, queries []int) []int {
    g := make(map[int]int)
    cnt := make([]int, n)
    for _, e := range edges {
        a, b := e[0] - 1, e[1] - 1
        a, b = min(a, b), max(a, b)
        cnt[a]++
        cnt[b]++
        g[a * n + b]++
    }

    ans := make([]int, len(queries))
    s := make([]int, n)
    copy(s, cnt)
    sort.Ints(s)
    for i, q := range queries {
        for j, x := range s {
            k := sort.Search(n, func(p int) bool {
                return p > j && x + s[p] > q
            })
            ans[i] += n - k 
        }

        for j, k := range g {
            a, b := j / n, j % n 
            if cnt[a] + cnt[b] > q && cnt[a] + cnt[b] - k <= q {
                ans[i]--
            }
        }
    }

    return ans
}