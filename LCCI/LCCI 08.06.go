func hanota(A []int, B []int, C []int) []int {

    var dfs func(i int, A, B, C *[]int)
    dfs = func(i int, A, B, C *[]int) {
        if i == 1 {
            *C = append(*C, (*A)[len(*A)-1])
            *A = (*A)[:len(*A)-1]
            return
        }

        dfs(i - 1, A, C, B)
        *C = append(*C, (*A)[len(*A)-1])
        *A = (*A)[:len(*A)-1]

        dfs(i - 1, B, A, C)
    }
    dfs(len(A), &A, &B, &C)

    return C
}