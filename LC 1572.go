func diagonalSum(mat [][]int) (ans int) {
    n := len(mat)

    for i := range mat {
        ans += mat[i][i] + mat[i][n-i-1]
    }
    
    if n & 1 == 1 {
        ans -= mat[n/2][n/2]
    }

    return 
}