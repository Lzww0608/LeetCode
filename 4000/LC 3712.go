func sumDivisibleByK(nums []int, k int) (ans int) {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }

    for a, b := range cnt {
        if b % k == 0 {
            ans += a * b
        } 
    }

    return 
}