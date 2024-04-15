func maxChunksToSorted(arr []int) (ans int) {
    mx := -1
    for i, x := range arr {
        mx = max(mx, x)
        if mx == i {
            ans++
        }
    }

    return 
}