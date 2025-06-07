func corpFlightBookings(bookings [][]int, n int) []int {
    ans := make([]int, n)
    dif := make([]int, n + 2)
    for _, booking := range bookings {
        dif[booking[0]] += booking[2]
        dif[booking[1] + 1] -= booking[2]
    }

    sum := 0
    for i := 1; i <= n; i++ {
        sum += dif[i]
        ans[i-1] = sum
    }

    return ans
}