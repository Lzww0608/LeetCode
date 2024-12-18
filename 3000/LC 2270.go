func waysToSplitArray(nums []int) (ans int) {
    sum := 0
    for _, x := range nums {
        sum += x 
    }

    cur := 0
    for _, x := range nums[:len(nums)-1] {
        cur += x 
        sum -= x 
        if cur >= sum {
            ans++
        }
    }

    return
}