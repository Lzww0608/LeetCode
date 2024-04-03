func maxOperations(nums []int, k int) (ans int) {
    m := map[int]int{}
    for _, x := range nums {
        if m[k-x] > 0 {
            ans++
            m[k-x]--
        } else {
            m[x]++
        }
    }

    return 
}