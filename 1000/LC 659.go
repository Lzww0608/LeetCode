func isPossible(nums []int) bool {
    m1, m2 := map[int]int{}, map[int]int{}
    for _, x := range nums {
        m1[x]++
    }

    for k := nums[0]; k <= nums[len(nums)-1]; k++ {
        for m1[k] > 0 && m2[k-1] > 0 {
            m1[k]--
            m2[k-1]--
            m2[k]++
        }
        for m1[k] > 0 && m1[k+1] > 0 && m1[k+2] > 0 {
            m1[k]--
            m1[k+1]--
            m1[k+2]--
            m2[k+2]++
            //fmt.Println(k+2)
        }
        if m1[k] > 0 {
            //fmt.Println(k, m1[k])
            return false
        }
    }
    return true
}
