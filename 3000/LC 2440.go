func componentValue(nums []int, edges [][]int) int {
    n := len(nums)
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    mx, sum := 0, 0
    for _, x := range nums {
        sum += x 
        mx = max(mx, x)
    }

    target := 0

    var dfs func(int, int) int 
    dfs = func(x, fa int) int {
        sum := nums[x]
        for _, y := range g[x] {
            if y != fa {
                t := dfs(y, x) 
                if t < 0 {
                    return -1
                }
                sum += t
            }
        }

        if sum > target {
            return -1
        }
        if sum == target {
            return 0
        }
        return sum
    }

    for x := sum / mx; x > 1; x-- {
        if sum % x == 0 {
            target = sum / x
            if dfs(0, -1) == 0 {
                return x - 1
            }
        }
    }

    return 0
}