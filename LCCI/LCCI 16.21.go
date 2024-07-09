func findSwapValues(a []int, b []int) []int {
    sumA, sumB := 0, 0
    for _, x := range a {
        sumA += x
    }
    for _, x := range b {
        sumB += x
    }

    if sumA < sumB {
        tmp := findSwapValues(b, a)
        if len(tmp) != 2 {
            return tmp
        }
        return []int{tmp[1], tmp[0]}
    }

    sort.Ints(a)
    sort.Ints(b)
    dif := sumA - sumB

    i, j := len(a) - 1, len(b) - 1 
    for i >= 0 && j >= 0 {
        x := a[i] - b[j]
        if x * 2 == dif {
            return []int{a[i], b[j]}
        } else if x * 2 > dif {
            i--
        }  else {
            j--
        }
    }

    return []int{}
}