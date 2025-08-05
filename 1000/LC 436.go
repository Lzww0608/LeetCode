func findRightInterval(intervals [][]int) []int {
    type pair struct {
        x, i int 
    }
    n := len(intervals)
    startSec := make([]pair, n)
    endSec := make([]pair, n)
    for i, v := range intervals {
        startSec[i] = pair{v[0], i}
        endSec[i] = pair{v[1], i}
    }

    sort.Slice(startSec, func(i, j int) bool {
        return startSec[i].x < startSec[j].x
    })
    sort.Slice(endSec, func(i, j int) bool {
        if endSec[i].x == endSec[j].x {
            return endSec[i].i < endSec[j].i
        }
        return endSec[i].x < endSec[j].x
    })

    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
    }
    for i, j := 0, 0; i < n && j < n; i++ {
        for j < n && startSec[j].x < endSec[i].x {
            j++
        }
        if j < n {
            ans[endSec[i].i] = startSec[j].i
        }
    }

    return ans
}