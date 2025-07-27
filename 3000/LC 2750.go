const MOD int = 1_000_000_007
func numberOfGoodSubarraySplits(nums []int) int {
    cnt := 0
    ans, pre := 1, 1
    for _, x := range nums {
        if x == 1 {
            if cnt == 0 {
                cnt = 1
            } else {
                ans = ans * (pre + 1) % MOD
            }
            pre = 0
        } else {
            pre++
        }
    }

    if cnt == 0 {
        return 0
    }

    return ans
}