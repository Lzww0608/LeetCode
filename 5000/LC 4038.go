func countSpecialIntegers(nums []int) (ans int) {
    a := []int{}
    m := make(map[int]int)
    for _, x := range nums {
        if len(a) == 0 || x != a[len(a) - 1] {
            a = append(a, x)
            m[x]++
        } 
    }

    for _, v := range m {
        if v == 1 {
            ans++
        }
    }

    return
}