func minimumDistance(nums []int) int {
    n := len(nums)
    ans := math.MaxInt32

    pos := make([][2]int, n)
    for i := range n {
        pos[i] = [2]int{-1, -1}
    }

    for i, x := range nums {
        x--
        if pos[x][1] != -1 {
            ans = min(ans, i - pos[x][0] + i - pos[x][1] + pos[x][0] - pos[x][1])
        }
        pos[x][0], pos[x][1] = i, pos[x][0]
        
    }

    if ans == math.MaxInt32 {
        return -1
    }
    return ans
}