func jump(nums []int) (ans int) {

    maxPos, end := 0, 0

    for i, x := range nums[:len(nums)-1] {
        maxPos = max(x + i, maxPos)
        if i == end {
            end = maxPos
            ans++
        }
    }

    return
}