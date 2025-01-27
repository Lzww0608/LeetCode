const N int = 30
const MOD int = 1_000_000_007
func maxSum(nums []int, k int) (ans int) {
    cnt := make([]int, N)
    for _, x := range nums {
        for i := 0; (1 << i) <= x; i++ {
            cnt[i] += (x >> i) & 1
        }
    }

    for ; k > 0; k-- {
        x := 0
        for i := range cnt {
            if cnt[i] > 0 {
                cnt[i]--
                x |= 1 << i
            }
        }
        ans += x * x
        ans %= MOD
    }
    return 
}