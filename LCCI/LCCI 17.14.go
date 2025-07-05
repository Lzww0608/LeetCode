func smallestK(arr []int, t int) []int {
    var quickSort func(int, int)
    quickSort = func(l, r int) {
        if l >= r {
            return
        }
        pivot := arr[l + rand.Intn(r - l + 1)]
        i, j, k := l, l, r + 1
        for i < k {
            if arr[i] > pivot {
                k--
                arr[i], arr[k] = arr[k], arr[i]
            } else if arr[i] < pivot {
                arr[i], arr[j] = arr[j], arr[i]
                i++
                j++
            } else {
                i++
            }
        }

        if i < t {
            quickSort(i, r)
        } else if t < j {
            quickSort(l, j)
        }
    }

    quickSort(0, len(arr)-1)
    return arr[:t]
}