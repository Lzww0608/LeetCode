/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 * 
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

func cleanRoom(robot *Robot) {
    vis := make(map[string]bool)
    dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    var dfs func(int, int, int)
    dfs = func(x, y, d int) {
        s := strconv.Itoa(x) + " " + strconv.Itoa(y)
        vis[s] = true
        robot.Clean()
        for i := 0; i < 4; i++ {
            pos := (d + i) % 4 
            next_x, next_y := x + dirs[pos][0], y + dirs[pos][1]
            t := strconv.Itoa(next_x) + " " + strconv.Itoa(next_y)
            if !vis[t] && robot.Move() {
                dfs(next_x, next_y, pos)
            }
            robot.TurnRight()
        }
        robot.TurnRight()
        robot.TurnRight()
        robot.Move()
        robot.TurnLeft()
        robot.TurnLeft()
    }
    dfs(0, 0, 0)

    return
}