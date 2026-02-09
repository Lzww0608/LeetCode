func minimumDistance(nums []int) int {
    n := len(nums)
    pos := make([][2]int, n + 1)
    for i := range pos {
        pos[i] = [2]int{-1, -1}
    }

    ans := math.MaxInt32
    for i, x := range nums {
        if pos[x][0] == -1 {
            pos[x][0] = i
        } else if pos[x][1] == -1 {
            pos[x][1] = i
        } else {
            ans = min(ans, i - pos[x][1] + i - pos[x][0] + pos[x][1] - pos[x][0])
            pos[x][0], pos[x][1] = pos[x][1], i
        }
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}