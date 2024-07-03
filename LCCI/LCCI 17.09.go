func getKthMagicNumber(k int) int {
    nums := make([]int, k)
    nums[0] = 1
    a, b, c := 0, 0, 0
    for i := 1; i < k; i++ {
        x, y, z := nums[a] * 3, nums[b] * 5, nums[c] * 7
        t := min(x, y, z)
        nums[i] = t
        if t == x {
            a++
        } 
        if t == y {
            b++
        }
        if t == z {
            c++
        }
    }

    return nums[k-1]
}