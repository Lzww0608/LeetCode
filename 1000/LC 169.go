func majorityElement(nums []int) int {
    res, cnt := -1, 0

    for _, x := range nums {
        if x == res {
            cnt++
        } else {
            cnt--
        }
        if cnt < 0 {
            res, cnt = x, 1
        }
    } 

    return res
}
