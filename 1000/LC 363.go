func maxSumSubmatrix(matrix [][]int, k int) int {
    ans := math.MinInt32
    m, n := len(matrix), len(matrix[0])

    for i := range matrix {
        sum := make([]int, n)
        for j := i; j < m; j++ {
            for c := 0; c < n; c++ {
                sum[c] += matrix[j][c]
            }
            sumSet := make([]int, 0, n + 1)
            sumSet = append(sumSet, 0)
            s := 0
            for _, v := range sum {
                s += v
                pos := sort.Search(len(sumSet), func(i int) bool {
                    return sumSet[i] >= s - k
                })
                if pos != len(sumSet) {
                    ans = max(ans, s - sumSet[pos])
                }
                if ans == k {
                    return ans
                }
                insert(&sumSet, s)
            }
        }
    }

    return ans
}

func insert(a *[]int, x int) {
    pos := sort.SearchInts(*a, x)
    *a = append(*a, 0)
    copy((*a)[pos+1:], (*a)[pos:])
    (*a)[pos] = x
    return 
}