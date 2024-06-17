func fullBloomFlowers(flowers [][]int, people []int) []int {
    n := len(flowers)
    start := make([]int, n)
    end := make([]int, n)
    for i, v := range flowers {
        start[i] = v[0]
        end[i] = v[1]
    }
    sort.Ints(start)
    sort.Ints(end)

    m := len(people)
    ans := make([]int, m)
    for i, x := range people {
        ans[i] = sort.SearchInts(start, x + 1) - sort.SearchInts(end, x)
    }

    return ans
}