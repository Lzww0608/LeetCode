func numberOfWeeks(milestones []int) int64 {
    mx := 0
    var sum int64 = 0
    for _, x := range milestones {
        mx = max(mx, x)
        sum += int64(x)
    }
    if sum >= int64(mx * 2) {
        return sum
    }
    return (sum - int64(mx)) * 2 + 1

}