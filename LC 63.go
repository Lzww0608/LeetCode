func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n := len(obstacleGrid[0])
    f := make([]int, n)
    f[0] = 1

    for i := range obstacleGrid {
        if obstacleGrid[i][0] == 1 {
            f[0] = 0
            continue
        } 
…}