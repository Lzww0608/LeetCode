func heightChecker(heights []int) (ans int) {
    a := make([]int, 101)
    for _, x := range heights {
        a[x]++
    }

    j := 0
    for i, x := range a {
        for ; x > 0; x-- {
            if heights[j] != i {
                ans++
            }
            j++
        }
    }
    return 
}
