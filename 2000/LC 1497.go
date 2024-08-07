func canArrange(arr []int, k int) bool {
    cnt := map[int]int{}
    for _, x := range arr {
        mod := ((x % k) + k) % k
        cnt[mod]++
    }
    
    for mod, count := range cnt {
        if mod == 0 || mod * 2 == k {
            if count%2 != 0 {
                return false
            }
        } else {
            if cnt[k-mod] != count {
                return false
            }
        }
    }
    
    return true
}
