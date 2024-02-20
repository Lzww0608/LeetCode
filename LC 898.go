func subarrayBitwiseORs(arr []int) int {
    m := map[int]int{}
    for i, x := range arr {
        m[x] = 1
        for j := i - 1; j >= 0; j-- {
            if arr[j] | x == arr[j] {
                break
            }//pruning
            arr[j] |= x
            m[arr[j]] = 1
        }
    }
    return len(m)
}