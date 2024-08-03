func maxPointsInsideSquare(points [][]int, s string) (ans int) {
    minD := [26]int{}
    for i := range minD {
        minD[i] = math.MaxInt32
    }

    min2 := math.MaxInt32
    for i, p := range points {
        d := max(abs(p[0]), abs(p[1]))
        c := s[i] - 'a'
        if d < minD[c] {
            min2 = min(min2, minD[c])
            minD[c] = d
        } else {
            min2 = min(min2, d)
        }
    }

    for _, d := range minD {
        if d < min2 {
            ans++
        }
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}