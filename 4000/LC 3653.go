const MOD int = 1_000_000_007
func xorAfterQueries(nums []int, queries [][]int) int {
    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3]
        for i := l; i <= r; i += k {
            nums[i] = (nums[i] * v) % MOD
        }
    }

    ans := 0
    for _, x := range nums {
        ans ^= x
    }

    return ans
}