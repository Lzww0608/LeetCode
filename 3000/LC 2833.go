func furthestDistanceFromOrigin(moves string) int {
    cur, cnt := 0, 0
    for _, c := range moves {
        if c == 'L' {
            cur--
        } else if c == 'R' {
            cur++
        } else {
            cnt++
        }
    }

    return abs(cur) + cnt
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    
    return x
}