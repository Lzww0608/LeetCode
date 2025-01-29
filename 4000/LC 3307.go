func kthCharacter(k int64, operations []int) byte {
    k--
    cnt := 0
    for i := 0; (1 << i) <= k; i++ {
        if (k >> i) & 1 == 1 {
            cnt += operations[i]
        }
    }

    return byte('a' + cnt % 26)
}