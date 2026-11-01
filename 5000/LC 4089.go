const N int = 27
func countMagicalPairs(s string) (ans int) {
    n := len(s)
    f := make([][N]int, n)
    for i := range n {
        vis := [N]bool{}
        cnt := 0
        for j := i; j < n; j++ {
            x := int(s[j] - 'a')
            if !vis[x] {
                cnt++
            }
            vis[x] = true
            f[i][cnt]++
        }
    }

    for i := range n - 1 {
        vis := [N]bool{}
        cnt := 0
        for j := i; j < n - 1; j++ {
            x := int(s[j] - 'a')
            if !vis[x] {
                cnt++
            }
            vis[x] = true
            ans += f[j + 1][cnt]
        }
    }

    return
}
