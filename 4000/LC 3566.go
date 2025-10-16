func checkEqualPartitions(nums []int, target int64) bool {
    n := len(nums)
    m := make(map[int]int)

    calc := func(mask int) int {
        if v, ok := m[mask]; ok {
            return v
        }

        res := 1
        for i := 0; (1 << i) <= mask; i++ {
            if mask & (1 << i) != 0 {
                res *= nums[i]
            }
        }
        m[mask] = res 
        return res
    }

    for i := 1; i + 1 < (1 << n) - 1; i++ {
        x := calc(i)
        if x != int(target) {
            continue
        }
        y := calc(((1 << n) - 1) &^ i)
        if y == int(target) {
            return true
        }
    } 

    return false 
}