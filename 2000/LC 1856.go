const MOD int = 1_000_000_007
func maxSumMinProduct(nums []int) (ans int) {
    n := len(nums)
    pre := make([]int, n + 1)
    for i, x := range nums {
        pre[i+1] = pre[i] + x
    }

    left := make([]int, n)
    st := []int{}
    for i, x := range nums {
        for len(st) > 0 && nums[st[len(st)-1]] >= x {
            st = st[:len(st)-1]
        }

        if len(st) == 0 {
            left[i] = -1
        } else {
            left[i] = st[len(st)-1]
        }
        st = append(st, i)
    }

    right := make([]int, n)
    st = []int{}
    for i := n - 1; i >= 0; i-- {
        for len(st) > 0 && nums[st[len(st)-1]] >= nums[i] {
            st = st[:len(st)-1]
        }

        if len(st) == 0 {
            right[i] = n
        } else {
            right[i] = st[len(st)-1]
        }
        st = append(st, i)
    }    

    for i, x := range nums {
        ans = max(ans, (pre[right[i]] - pre[left[i]+1]) *x )
    }

    return ans % MOD
}