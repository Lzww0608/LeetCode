func maximizeGreatness(nums []int) (ans int) {
    m := make(map[int]int)
    for _, x := range nums {
        m[x]++
        ans = max(ans, m[x])
    }
    return len(nums) - ans
}
