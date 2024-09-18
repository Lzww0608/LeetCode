
import (
	"math"
	"math/bits"
)
func minimumIncompatibility(nums []int, k int) int {
    m := map[int]int{}
    for _, x := range nums {
        m[x]++
        if m[x] > k {
            return -1
        }
    }

    clear(m)
    n := len(nums)
    sz := n / k
    sort.Ints(nums)

    var dfs func(int, int) int
    dfs = func(left, pre int) int {
        if left == 0 {
            return 0
        }

        if bits.OnesCount(uint(left)) % sz == 0 {
            lb := left & (-left)
            return dfs(left ^ lb, ctz(lb))
        }

        if m[(left << 4) | pre] != 0 {
            return m[(left << 4) | pre]
        }

        res := math.MaxInt32
        last := nums[pre]

        for i := pre + 1; i < n; i++ {
            if (left >> i) & 1 == 1 && nums[i] != last {
                last = nums[i]
                res = min(res, last - nums[pre] + dfs(left ^ (1 << i), i))
            }
        }

        m[(left << 4) | pre] = res
        return res
    }


    return dfs(1 << n - 2, 0)
}

func ctz(x int) (res int) {
    for x & 1 == 0 {
        res++
        x >>= 1
    }

    return 
}