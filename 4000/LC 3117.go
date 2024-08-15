func minimumValueSum(nums []int, andValues []int) int {
    n, m := len(nums), len(andValues)
    type tuple struct {
        i, j, and int
    }

    memo := map[tuple]int{}

    var dfs func(int, int, int) int
    dfs = func(i, j, and int) int {
        if n - i < m - j {
            return math.MaxInt32
        }

        if j == m {
            if i == n {
                return 0
            }
            return math.MaxInt32
        }

        and &= nums[i]
        p := tuple{i, j, and}
        if res, ok := memo[p]; ok {
            return res
        }
        res := dfs(i + 1, j, and)
        if and == andValues[j] {
            res = min(res, dfs(i + 1, j + 1, -1) + nums[i])
        }
        memo[p] = res
        return res
    }

    ans := dfs(0, 0, -1)
    if ans == math.MaxInt32 {
        return -1
    }

    return ans
}