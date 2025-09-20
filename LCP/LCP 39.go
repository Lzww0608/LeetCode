func minimumSwitchingTimes(source [][]int, target [][]int) (ans int) {
    cnt := make(map[int]int)
    for _, v := range source {
        for _, x := range v {
            cnt[x]++
        }
    }

    for _, v := range target {
        for _, x := range v {
            cnt[x]--
        }
    }

    for _, x := range cnt {
        ans += abs(x)
    }

    return ans / 2
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}