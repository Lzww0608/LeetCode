const N int = 101
func twoOutOfThree(a []int, b []int, c []int) (ans []int) {
    cnt := [N]int{}
    for _, x := range a {
        cnt[x] |= 1
    }
    for _, x := range b {
        cnt[x] |= 2
    }
    for _, x := range c {
        cnt[x] |= 4
    }

    for k, x := range cnt {
        if x & (x - 1) != 0 {
            ans = append(ans, k)
        }
    }

    return
}