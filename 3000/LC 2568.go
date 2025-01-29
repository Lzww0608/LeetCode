func minImpossibleOR(nums []int) int {
    m := make(map[int]bool)
    for _, x := range nums {
        m[x] = true
    }

    ans := 1
    for m[ans] {
        ans <<= 1
    }

    return ans
}