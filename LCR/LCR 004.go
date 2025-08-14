const N int = 32
func singleNumber(nums []int) int {
    cnt := make([]int32, N)
    for _, x := range nums {
        for i := 0; i < 32; i++ {
            if x & (1 << i) != 0 {
                cnt[i]++
            }
        }
    }

    var ans int32
    for i, x := range cnt {
        if x % 3 == 1 {
            ans |= 1 << i
        }
    }

    return int(ans)
}