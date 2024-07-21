func twoSum(a []int, target int) []int {
    l, r := 0, len(a) - 1
    for {
        x := a[l] + a[r]
        if x == target {
            return []int{l, r}
        } else if x > target {
            r--
        } else {
            l++
        }
    }

    return nil
}