func findLucky(arr []int) int {
    cnt := make([]int, 505)
    for _, x := range arr {
        cnt[x]++
    }

    for i := len(cnt) - 1; i > 0; i-- {
        if cnt[i] == i {
            return i
        }
    }

    return -1
}