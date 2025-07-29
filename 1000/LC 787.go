
import "math"
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    dis := make([]int, n)
    for i := range dis {
        dis[i] = math.MaxInt32
    }
    dis[src] = 0

    for i := 0; i <= k; i++ {
        tmp := make([]int, n)
        copy(tmp, dis)
        for _, v := range flights {
            if tmp[v[0]] != math.MaxInt32 {
                dis[v[1]] = min(dis[v[1]], tmp[v[0]] + v[2])
            }
        }
    }

    if dis[dst] >= math.MaxInt32 {
        return -1
    }
    return dis[dst]
}