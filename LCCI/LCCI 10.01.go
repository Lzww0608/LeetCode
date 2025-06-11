func merge(A []int, m int, B []int, n int)  {
    i, j, k := m - 1, n - 1, m + n - 1 
    for i >= 0 || j >= 0 {
        if j < 0 || i >= 0 && A[i] >= B[j] {
            A[k] = A[i]
            i--
        } else {
            A[k] = B[j]
            j--
        }
        k--
    }
}