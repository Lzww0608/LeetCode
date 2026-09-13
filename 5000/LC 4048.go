func countSpecialIntegers(nums []int) (ans int) {
    pos := make(map[int][]int)
    for i, x := range nums {
        pos[x] = append(pos[x], i)
    } 

    for _, v := range pos {
        if len(v) != 3 {
            continue
        }
        if v[2] - v[1] == v[1] - v[0] {
            ans++
        }
    }

    return
}