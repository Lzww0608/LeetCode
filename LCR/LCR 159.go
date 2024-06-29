func inventoryManagement(stock []int, cnt int) []int {
    n := len(stock)

    var quickSort func(int, int)
    quickSort = func(l, r int) {
        if l >= r {
            return
        }
        pivot := stock[l + rand.Intn(r - l + 1)]
        i, j, k := l, l, r + 1
        for i < k {
            if stock[i] < pivot {
                stock[i], stock[j] = stock[j], stock[i]
                i++
                j++
            } else if stock[i] > pivot {
                k--
                stock[i], stock[k] = stock[k], stock[i]
            } else {
                i++
            }
        }
        if k < cnt {
            quickSort(k, r)
        } else if j >= cnt {
            quickSort(l, j - 1)
        } 
        return
    }
    quickSort(0, n - 1)

    return stock[:cnt]
}
