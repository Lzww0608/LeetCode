const MOD int = 1_000_000_007
func sumOfPowers(nums []int, k int) (ans int) {
    sort.Ints(nums)
    n := len(nums)

    var f func([]int, int) []int
    f = func(a []int, lower_diff int) []int {
        n := len(a)
        f := make([][]int, n)
        for i := range f {
            f[i] = make([]int, k)
        }

        f[0][1] = 1

        for i := 1; i < n; i++ {
            for j := 0; j < i; j++ {
                if a[i] - a[j] < lower_diff {
                    break
                }
                for v := 0; v < k - 1; v++ {
                    f[i][v+1] += f[j][v]
                    f[j][v+1] %= MOD
                }
            }
        }

        ans := make([]int, k)
        for i := 0; i < n; i++ {
            for j := 0; j < k; j++ {
                ans[j] += f[i][j]
            }
        }

        return ans
    }

    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            v := nums[i] - nums[j]
            vis1 := []int{}
            for idx := j; idx >= 0; idx-- {
                vis1 = append(vis1, nums[j] - nums[idx])
            }

            dp1 := f(vis1, v + 1)

            vis2 := []int{}
            for idx := i; idx < n; idx++ {
                vis2 = append(vis2, nums[idx] - nums[i])
            }

            dp2 := f(vis2, v)

            for x := 1; x < k; x++ {
                ans += dp1[x] * dp2[k-x] * v
                ans %= MOD
            }
        }
    }

    return
}