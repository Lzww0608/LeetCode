func numFriendRequests(ages []int) int {
    check := func(x, y int) bool {
        if y <= x / 2 + 7 || y > x || (y > 100 && x < 100) {
            return false
        }
        return true
    }
    ans, n := 0, len(ages)
    sort.Ints(ages)
    for i := range ages {
        j := i + 1
        for j < n && check(ages[i], ages[j]) {
            ans++
            j++
        }
        j = i - 1
        for j >= 0 && check(ages[i], ages[j]) {
            ans++
            j--
        }
    }
    return ans
}