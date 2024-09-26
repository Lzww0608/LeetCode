import (
    "sort"
)
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
    type pair struct {
        x, y int
    }
    n := len(nums1)
    a := make([]pair, n)
    for i, x := range nums1 {
        a[i] = pair{x, nums2[i]}
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i].x > a[j].x
    })

    qid := make([]int, len(queries))
    for i := range qid {
        qid[i] = i
    }
    sort.Slice(qid, func(i, j int) bool {
        return queries[qid[i]][0] > queries[qid[j]][0]
    })

    ans := make([]int, len(queries))
    type data struct {
        y, s int
    }
    st := []data{}
    j := 0
    for _, i := range qid {
        x, y := queries[i][0], queries[i][1]
        for ; j < n && a[j].x >= x; j++ {
            for len(st) > 0 && st[len(st)-1].s <= a[j].x + a[j].y {
                st = st[:len(st)-1]
            }
            if len(st) == 0 || st[len(st)-1].y < a[j].y {
                st = append(st, data{a[j].y, a[j].x + a[j].y})
            }
        }
        pos := sort.Search(len(st), func(i int) bool {
            return st[i].y >= y
        })
        if pos == len(st) {
            ans[i] = -1
        } else {
            ans[i] = st[pos].s
        }
    }

    return ans
}