func minimumRounds(tasks []int) int {
    m := map[int]int{}
    for _, x := range tasks {
        m[x]++
    }

    ans := 0
    for _, v := range m {
        if v == 1 {
            return -1
        }
        ans += (v + 2) / 3

    }

    return ans
}
