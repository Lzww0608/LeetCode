
import "sort"
func maximumWhiteTiles(tiles [][]int, carpetLen int) (ans int) {
    n := len(tiles)
    sort.Slice(tiles, func(i, j int) bool {
        return tiles[i][0] < tiles[j][0]
    })

    pre := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        pre[i] = pre[i-1] + (tiles[i-1][1] - tiles[i-1][0] + 1)
    }
    if carpetLen >= tiles[n-1][1] {
        return pre[n]
    }
    for i := 0; i < n; i++ {
        r := tiles[i][0] + carpetLen
        if r > tiles[n-1][1] {
            ans = max(ans, pre[n] - pre[i])
            break
        }
        j := sort.Search(len(tiles), func(k int) bool {
            return tiles[k][1] >= r
        })
        cur := pre[j+1] - pre[i] - min(tiles[j][1] - r, tiles[j][1] - tiles[j][0]) - 1
        //fmt.Println(i, cur)
        ans = max(ans, cur)
    }
    
    return 
}