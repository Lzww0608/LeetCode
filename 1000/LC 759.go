

import "sort"
/**
 * Definition for an Interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */

func employeeFreeTime(schedule [][]*Interval) (ans []*Interval) {
    a := []*Interval{}
    for i := range schedule {
        a = append(a, schedule[i]...)
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i].Start < a[j].Start
    })

    r := a[0].Start
    for _, v := range a {
        if r < v.Start {
            ans = append(ans, &Interval{r, v.Start})
        }
        r = max(r, v.End)
    }

    return
}