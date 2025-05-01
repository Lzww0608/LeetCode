/**
 * // This is Sea's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Sea struct {
 *     func hasShips(topRight, bottomLeft []int) bool {}
 * }
 */

func countShips(sea Sea, topRight, bottomLeft []int) int {
    
    var dfs func([]int, []int, bool) int
    dfs = func(a, b []int, f bool) (sum int) {
        if a[0] < b[0] || a[1] < b[1] {
            return 0
        }
        
        if !f && !sea.hasShips(a, b) {
            return 0
        }

        if a[0] == b[0] && a[1] == b[1] {
            return 1
        }

        if a[0] == b[0] {
            mid := (a[1] + b[1]) / 2
            sum = dfs(a, []int{a[0], mid + 1}, false)
            sum += dfs([]int{a[0], mid}, b, sum == 0)  
        } else {
            mid := (a[0] + b[0]) / 2 
            sum = dfs(a, []int{mid + 1, b[1]}, false)
            sum += dfs([]int{mid, a[1]}, b, sum == 0)
        }

        return
    }

    return dfs(topRight, bottomLeft, false)
}