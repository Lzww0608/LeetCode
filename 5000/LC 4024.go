func nearestDrone(drones [][]int, target []int) int {
    ans, mn := -1, math.MaxInt32
    for i, v := range drones {
        x, y, d := v[0], v[1], v[2]
        m := abs(x - target[0]) + abs(y - target[1])
        if m <= d {
            if m < mn {
                ans = i
                mn = m
            }
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}