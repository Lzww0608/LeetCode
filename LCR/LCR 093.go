func lenLongestFibSubseq(arr []int) (ans int) {
    n := len(arr)
    f := make([][]int, n)
    m := map[int]int{}

    for i, x := range arr {
        f[i] = make([]int, n)
        for j := i - 1; j >= 0 && arr[j] >= x / 2; j-- {
            if v, ok := m[x-arr[j]]; ok && v < j {
                f[j][i] = max(3, f[v][j] + 1)
            }
            ans = max(ans, f[j][i])
        }  
        m[x] = i
    }

    return 
}