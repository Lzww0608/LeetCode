func kClosest(points [][]int, q int) [][]int {
    f := func(i int) int {
        x, y := points[i][0], points[i][1]
        return x * x + y * y
    }

    var quickSort func(int, int)
    quickSort = func(l, r int) {
        if l >= r {
            return 
        }

        pivot := f(l + rand.Intn(r - l + 1))
        i, j, k := l, l, r + 1
        for i < k {
            t := f(i)
            if t > pivot {
                k--
                points[i], points[k] = points[k], points[i]
            } else if t < pivot {
                points[i], points[j] = points[j], points[i]
                i++
                j++
            } else {
                i++
            }
        }

        if q < j {
            quickSort(l, j - 1)
        } else if q > i {
            quickSort(i, r)
        }
        return
    }
    quickSort(0, len(points) - 1)

    return points[:q]
}