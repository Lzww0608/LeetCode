func findLeastNumOfUniqueInts(arr []int, k int) (ans int) {
    m := map[int]int{}
    for _, x := range arr {
        m[x]++
    }
    freq := make([]int, len(arr) + 1)
    ans += len(m)
    for _, v := range m {
        freq[v]++
    }
    for i, x := range freq {
        if x == 0 {
            continue
        }
        if k < i {
            break
        }
        t := min(x, k / i)
        k -= t * i
        ans -= t
    }

    return

}