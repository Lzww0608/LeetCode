import (
    "sort"
    "math"
)
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
    sort.Ints(arr2)
    n := len(arr1)
    m := make([]map[int]int, n)
    for i := range m {
        m[i] = make(map[int]int)
    }

    var dfs func(int, int) int 
    dfs = func(i, pre int) (res int) {
        if i < 0 {
            return 0
        }

        if v, ok := m[i][pre]; ok {
            return v
        } else {
            defer func() {m[i][pre] = res} ()
        }

        if arr1[i] < pre {
            res = dfs(i - 1, arr1[i])
        } else {
            res = math.MaxInt32 / 2
        }

        pos := sort.SearchInts(arr2, pre)
        if pos > 0 {
            res = min(res, dfs(i - 1, arr2[pos - 1]) + 1)
        }

        return
    }

    ans := dfs(n - 1, math.MaxInt32)
    if ans >= math.MaxInt32 / 2 {
        return -1
    }

    return ans
}