func firstUniqueFreq(nums []int) int {
    a := make(map[int]int)
    b := make(map[int]int)

    for _, x := range nums {
        a[x]++
    }
    for _, v := range a {
        b[v]++
    }

    for _, x := range nums {
        y := a[x]
        if b[y] == 1 {
            return x
        }
    }

    return -1
}