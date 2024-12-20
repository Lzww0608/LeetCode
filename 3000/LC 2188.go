import (
    "math"
)
func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
    minCost := make([]int, numLaps + 1)
    for i := range minCost {
        minCost[i] = math.MaxInt32
    }
    for _, v := range tires {
        cur, sum := v[0], v[0]
        minCost[1] = min(minCost[1], sum)
        for j := 2; j <= numLaps && sum <= math.MaxInt32; j++ {
            cur *= v[1]
            sum += cur
            minCost[j] = min(minCost[j], sum)
        }
    }

    f := make([]int, numLaps + 1)
    f[0] = -changeTime 
    for i := 1; i <= numLaps; i++ {
        f[i] = math.MaxInt32
        for j := 1; j <= min(i, len(minCost)-1) && minCost[j] < math.MaxInt32; j++ {
            f[i] = min(f[i], f[i-j] + minCost[j])
        } 
        f[i] += changeTime
    }

    return f[numLaps]
}