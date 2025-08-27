func fileCombination(target int) (ans [][]int) {
    sum := 0
    l, r := 1, 1
    for r <= target {
        sum += r 
        for sum > target {
            sum -= l
            l++
        }
        if sum == target && l < r {
            tmp := []int{}
            for i := l; i <= r; i++ {
                tmp = append(tmp, i)
            }
            ans = append(ans, tmp)
        } 
        r++
    }

    return 
}
