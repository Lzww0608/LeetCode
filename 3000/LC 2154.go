func findFinalValue(nums []int, original int) int {
    mask := 0
    for _, x := range nums {
        if x % original == 0 {
            k := x / original
            if k & (k - 1) == 0 {
                mask |= k 
            }
        }
    }

    mask = ^mask
    return original * (mask & -mask)
}