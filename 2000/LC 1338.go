func minSetSize(arr []int) (ans int) {
    n := len(arr)
    freq := map[int]int{}
    for _, x := range arr {
        freq[x]++
    }

    a := []int{}
    for _, x := range freq {
        a = append(a, x)
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })
    sum := 0
    for _, x := range a {
        sum += x
        ans++
        if sum >= n / 2 {
            break
        }
    }

    return
}