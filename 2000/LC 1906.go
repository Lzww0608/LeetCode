const N int = 101
func minDifference(nums []int, queries [][]int) []int {
    n := len(queries)
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    sort.Slice(id, func(i, j int) bool {
        return queries[id[i]][0] < queries[id[j]][0]
    })

    ans := make([]int, n)
    a := make([][]int, N)
    for i, x := range nums {
        a[x] = append(a[x], i)
    }

    pos := 0
    for _, i := range id {
        l, r := queries[i][0], queries[i][1]
        for pos < l {
            x := nums[pos]
            a[x] = a[x][1:]
            pos++
        }
        cur := N
        pre := -N
        for j := 1; j < N; j++ {
            if len(a[j]) > 0 {
                pos := sort.SearchInts(a[j], r + 1)
                if pos > 0 {
                    cur = min(cur, j - pre)
                    pre = j
                }
            }
        }

        if cur >= N {
            ans[i] = -1
        } else {
            ans[i] = cur
        }
    }

    return ans
}