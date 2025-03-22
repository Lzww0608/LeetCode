func rowAndMaximumOnes(mat [][]int) []int {
    ans := []int{0, 0}
    for i := 0; i < len(mat); i++ {
        cnt := 0
        for _, x := range mat[i] {
            if x == 1 {
                cnt++
            }
        }

        if cnt > ans[1] {
            ans = []int{i, cnt}
        }
    }

    return ans
}