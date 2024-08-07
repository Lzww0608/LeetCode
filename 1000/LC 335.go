func isSelfCrossing(distance []int) bool {
    n := len(distance)
    if n < 4 {
        return false
    }

    for i := 3; i < n; i++ {
        if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
            return true 
        } else if i >= 4 && distance[i-1] == distance[i-3] && distance[i] + distance[i-4] >= distance[i-2] {
            return true
        } else if (i >= 5 && distance[i] + distance[i-4] >= distance[i-2] && distance[i-2] > distance[i-4] &&
         distance[i-1] + distance[i-5] >= distance[i-3] && distance[i-3] >= distance[i-1]) {
                return true
            }
    }

    return false
}