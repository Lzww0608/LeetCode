func canTransform(start string, end string) bool {
    i, j, n := 0, 0, len(start)
    for i < n || j < n {
        for i < n && start[i] == 'X' {
            i++
        }
        for j < n && end[j] == 'X' {
            j++
        }
        if i == n || j == n {
            return i == j
        }
        if start[i] != end[j] {
            return false
        }
        if start[i] == 'R' && i > j {
            return false
        } else if start[i] == 'L' && i < j {
            return false
        }
        i, j = i + 1, j + 1
    }
    return i == j
}
