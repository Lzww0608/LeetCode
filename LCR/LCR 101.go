func canPartition(nums []int) bool {
    sum := 0
    for _, x := range nums {
        sum += x
    }

    if sum & 1 == 1 {
        return false
    }

    target := sum >> 1
    //fmt.Println(sum, target)
    f := make([]bool, target + 1)
    f[0] = true
    //sort.Ints(nums)
    for _, x := range nums {
        for k := target; k >= x; k-- {
            f[k] = f[k] || f[k-x]
        }
    }

    return f[target]
}