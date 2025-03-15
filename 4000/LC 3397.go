func maxDistinctElements(nums []int, k int) (ans int) {
    sort.Ints(nums)
    pre := math.MinInt32 

    for _, x := range nums {
        if x - k > pre {
            pre = x - k 
            ans++
        } else if x + k > pre {
            pre++
            ans++
        }
    }
    return
}