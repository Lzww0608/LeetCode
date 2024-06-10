func largestSquareArea(bottomLeft [][]int, topRight [][]int) (ans int64) {
    n := len(bottomLeft)
    for i := range bottomLeft {
        x1, y1, x2, y2 := bottomLeft[i][0], bottomLeft[i][1], topRight[i][0], topRight[i][1]
        for j := i + 1; j < n; j++ {
            x3, y3, x4, y4 := bottomLeft[j][0], bottomLeft[j][1], topRight[j][0], topRight[j][1]
            if x2 > x3 && y2 > y3 && x1 < x4 && y1 < y4{
                width := min(x2, x4) - max(x1, x3)
                height := min(y2, y4) - max(y1, y3)
                d := int64(min(width, height))
                ans = max(ans, d * d)
            } 
        }
    }
    return
}
