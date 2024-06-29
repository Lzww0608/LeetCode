func reverseBits(num int) (ans int) {
    pre, cnt := -1, 0
    for i := 0; i < 32; i++ {
        if (num >> i) & 1 == 1 {
            cnt++
        } else {
            pre = cnt
            cnt = 0
        }
        ans = max(ans, pre + 1 + cnt)
    }

    return
}