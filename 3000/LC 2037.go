func minMovesToSeat(seats []int, students []int) (ans int) {
    n := len(seats)
    sort.Ints(seats)
    sort.Ints(students)
    for i := range n {
        ans += abs(seats[i] - students[i])
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}