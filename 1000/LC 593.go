
import "reflect"
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
    type pair struct {
        x, i int 
    }

    calc := func(a, b []int) int {
        return (a[0] - b[0]) * (a[0] - b[0]) + (a[1] - b[1]) * (a[1] - b[1])
    }

    d := []pair {
        {calc(p1, p2), 2},
        {calc(p1, p3), 3},
        {calc(p1, p4), 4},
    }
    sort.Slice(d, func(i, j int) bool {
        return d[i].x < d[j].x
    })

    if d[0].x + d[1].x != d[2].x || reflect.DeepEqual(p1, p2) ||
        reflect.DeepEqual(p2, p3) || reflect.DeepEqual(p1, p3) {
        return false
    }

    var d1, d2 int 
    if d[2].i == 2 {
        d1, d2 = calc(p2, p3), calc(p2, p4)
    } else if d[2].i == 3 {
        d1, d2 = calc(p3, p2), calc(p3, p4)
    } else {
        d1, d2 = calc(p3, p4), calc(p2, p4)
    }

    return d1 == d2 && d1 == d[0].x
}