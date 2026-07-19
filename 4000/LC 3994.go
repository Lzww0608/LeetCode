const MOD int = 1_000_000_007
func minAdjacentSwaps(nums []int, a int, b int) int {
    n := len(nums)
    for i, x := range nums {
        if x < a {
            nums[i] = 0
        } else if x > b {
            nums[i] = 2
        } else {
            nums[i] = 1
        }
    }

    var merge func(l, r int) int
    merge = func(l, r int) int {
        if l >= r {
            return 0
        }

        mid := l + ((r - l) >> 1)
        cnt := merge(l, mid) + merge(mid + 1, r)
        for i, j := l,  mid + 1; i <= mid; i++ {
            for j <= r && nums[j] < nums[i] {
                j++
            } 

            cnt += j - mid - 1
        }

        tmp := make([]int, r - l + 1)
        i, j, k := l, mid + 1, 0
        for i <= mid || j <= r {
            if j > r || i <= mid && nums[i] <= nums[j] {
                tmp[k] = nums[i]
                i++
            } else {
                tmp[k] = nums[j]
                j++
            }
            k++
        }

        copy(nums[l:r+1], tmp)
        return cnt % MOD
    }

    return merge(0, n - 1)
}