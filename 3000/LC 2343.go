func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
    n := len(nums[0])
    a := make([][]int, n + 1)
    for i := range nums {
        a[0] = append(a[0], i)
    }

    for i := 1; i <= n; i++ {
        cur := make([][]int, 10)
        for _, x := range a[i-1] {
            idx := int(nums[x][n-i] - '0')
            cur[idx] = append(cur[idx], x)
        }
        for j := 0; j < 10; j++ {
            for _, x := range cur[j] {
                a[i] = append(a[i], x)
            }
        }
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        ans[i] = a[q[1]][q[0]-1]
    }

    return ans
}