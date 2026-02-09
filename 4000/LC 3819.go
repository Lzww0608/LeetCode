func rotateElements(nums []int, k int) []int {
    type pair struct {
        x, i int
    }
    pos := []pair{}
    for i, x := range nums {
        if x >= 0 {
            pos = append(pos, pair{x, i})
        }
    }

    n := len(pos)
    if n != 0 {
        k %= n 
    }
    for i, v := range pos {
        j := (i - k + n) % n 
        nums[pos[j].i] = v.x
    }

    return nums
}