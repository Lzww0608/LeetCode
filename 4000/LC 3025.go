
import "math"
func numberOfPairs(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][0] == points[j][0] {
            return points[i][1] > points[j][1]
        }
        return points[i][0] < points[j][0]
    })

    ans := 0
    for i, v := range points {
        y := math.MinInt32
        for _, t := range points[i + 1:] {
            if t[1] > y && t[1] <= v[1] {
                y = t[1]
                ans++
            } 
        }
    }

    return ans
}