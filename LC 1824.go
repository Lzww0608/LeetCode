func minSideJumps(obstacles []int) int {
    n := len(obstacles) - 1
    f := make([]int, 3)
    f[0], f[1], f[2] = 1, 0, 1

    for i := 1; i < n; i++ {
        if obstacles[i] == 0 {
            f[0], f[1], f[2] = min(f[0], f[1] + 1, f[2] + 1), min(f[1], f[0] + 1, f[2] + 1), min(f[2], f[1] + 1, f[0] + 1)

        } else {
            f[obstacles[i]-1] = math.MaxInt32 / 2
            if obstacles[i] == 1 {
                f[1], f[2] = min(f[1], f[0] + 1, f[2] + 1), min(f[2], f[1] + 1, f[0] + 1)
            } else if obstacles[i] == 2 {
                f[0], f[2] = min(f[0], f[1] + 1, f[2] + 1), min(f[2], f[1] + 1, f[0] + 1)
            } else {
                f[0], f[1] = min(f[0], f[1] + 1, f[2] + 1), min(f[1], f[2] + 1, f[0] + 1)
            }
        } 
    }

    return min(f[0], f[1], f[2])
}