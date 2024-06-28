func maxDistToClosest(seats []int) int {
    n := len(seats)
    ans := 0
    l, r := -1, 0
    for r < n {
        for r < n && seats[r] == 0 {
            r++
        }
        if l == -1 || r == n{
            ans = max(ans, r - l - 1)
        } else {
            ans = max(ans, (r - l) / 2)
        }
        l, r = r, r + 1
    }
    //ans = max(ans, r - l - 1)
    return ans
}
