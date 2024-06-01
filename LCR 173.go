func takeAttendance(records []int) int {
    n := len(records)
    l, r := 0, n
    for l < r {
        mid := l + ((r - l) >> 1)
        if records[mid] > mid {
            r = mid 
        } else {
            l = mid + 1
        }
    }

    return l
}