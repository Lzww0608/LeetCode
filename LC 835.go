type pair struct {
    x, y int
}
func largestOverlap(img1 [][]int, img2 [][]int) int {
    n := len(img1)
    ans := 0
    m := map[pair]int{}
    rec := map[pair]bool{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if img2[i][j] == 1 {
                rec[pair{i, j}] = true
            }
        }
    }

    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if img1[i][j] == 1 {
                for v := range rec {
                    m[pair{v.x-i, v.y-j}]++
                }
            }
        }
    }
    for _, v := range m {
        ans = max(ans, v)
    }
    return ans
}