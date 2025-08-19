func numSubarraysWithSum(nums []int, goal int) (ans int) {
    m := make(map[int]int)
    m[0] = 1
    pre := 0
    for _, x := range nums {
        if x == 1 {
            pre++
        }
        ans += m[pre - goal]
        m[pre]++
    }

    return
}