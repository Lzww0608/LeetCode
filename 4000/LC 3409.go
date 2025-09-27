const N int = 301
func longestSubsequence(nums []int) (ans int) {
    //f[i][j] means the maximum length ends with i and the last dif >= j
    f := [N][N]int{}
    for _, x := range nums {
        cur := 1
        for y := N - 1; y >= 0; y-- {
            if x - y >= 0 {
                cur = max(cur, f[x - y][y] + 1)
            }
            if x + y < N {
                cur = max(cur, f[x + y][y] + 1)
            }
            f[x][y] = cur 
            ans = max(ans, cur)
        }
    }
    
    return 
}