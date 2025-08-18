func findBestValue(arr []int, target int) int {
    sort.Ints(arr)
    n := len(arr)
    pre := 0

    for i, x := range arr {
        cur := pre + (n - i) * x
        if cur >= target {
            t := (target - pre) / (n - i)
            l, r := pre + t * (n - i), pre + (t + 1) * (n - i)
            if target - l <= r - target {
                return t 
            }
            return t + 1
        }
        pre += x
    }

    return arr[n - 1]
}