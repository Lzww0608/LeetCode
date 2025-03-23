func minimumPartition(s string, k int) (ans int) {
    cur := 0
    for i := range s {
        x := int(s[i] - '0')
        if x > k {
            return -1
        }
        cur = cur * 10 + x 
        if cur > k {
            ans++
            cur = x 
        }
    }

    return ans + 1
}