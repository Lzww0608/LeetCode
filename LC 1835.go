func getXORSum(arr1 []int, arr2 []int) int {
    a, b := 0, 0
    ans := 0

    for _, x := range arr1 {
        a ^= x
    }
    for _, x := range arr2 {
        ans ^= x & a
        b ^= x
    }

    return ans
}