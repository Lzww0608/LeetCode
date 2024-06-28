func pairSums(nums []int, target int) (ans [][]int) {
    m := map[int]int{}
    for _, x := range nums {
        if m[target-x] > 0 {
            m[target-x]--
            ans = append(ans, []int{target-x, x})
        } else {
            m[x]++
        }
    }

    return
}