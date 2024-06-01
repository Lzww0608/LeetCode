func nearestValidPoint(x int, y int, points [][]int) int {
    ans, mn:= -1, math.MaxInt32

    for i, p := range points {
        a, b := p[0], p[1]
        if a != x && b != y {
            continue
        }  
        dis := abs(a - x) + abs(b - y)
        if dis < mn {
            mn = dis
            ans = i
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
