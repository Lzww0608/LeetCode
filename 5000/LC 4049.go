func countSpecialIntegers(nums []int) (ans int) {
    pos := make(map[int][]int)
    for i, x := range nums {
        pos[x] = append(pos[x], i)
    } 

next:
    for _, v := range pos {
        if len(v) < 3 {
            continue
        }
        for i := 2; i < len(v); i++ {
            if v[i] - v[i - 1] != v[1] - v[0] {
                continue next
            }
        }

        ans++
    }

    return
}