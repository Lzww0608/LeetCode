func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {

    m := map[int]int{}

    type pair struct {
        x, y int
    }
    a := []pair{}
    for i := range values {
        a = append(a, pair{values[i], labels[i]})
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i].x >= a[j].x
    })

    ans := 0
    for i := 0; numWanted > 0 && i < len(a); i++ {
        if m[a[i].y] < useLimit {
            numWanted--
            m[a[i].y]++
            ans += a[i].x
        }
        
    }

    return ans
}
