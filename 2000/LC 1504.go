func numSubmat(mat [][]int) (ans int) {
    n := len(mat[0])
    pre := make([]int, n)

    for i := range mat {
        for j := range mat[i] {
            if mat[i][j] == 1 {
                pre[j]++
            } else {
                pre[j] = 0
            }
        }

        st := []int{-1}
        sum := 0
        for j := range mat[i] {
            for len(st) > 1 && pre[j] < pre[st[len(st)-1]] {
                k := st[len(st)-1]
                st = st[:len(st)-1]
                sum -= (k - st[len(st)-1]) * pre[k] 
            }
            sum += (j - st[len(st)-1]) * pre[j]
            ans += sum
            st = append(st, j)
        }
    }
    return
}
