func unequalTriplets(nums []int) (ans int) {
    cnt := make(map[int]int)
    a, b := len(nums), 0
    for _, x := range nums {
        cnt[x]++
    }

    for _, x := range cnt {
        a -= x 
        ans += a * b * x 
        b += x
    }

    return
}