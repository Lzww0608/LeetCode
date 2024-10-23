func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
    n := len(nums1)
    a := make([]int, n)
    for i := range a {
        a[i] = nums1[i] - nums2[i]
    }

    return int64(mergeSort(a, 0, n - 1, diff))
}

func mergeSort(a []int, l, r, diff int) (ans int) {
    if l >= r {
        return
    }

    mid := l + ((r - l) >> 1)
    ans += mergeSort(a, l, mid, diff) + mergeSort(a, mid + 1, r, diff)
    
    for i, j := mid, r; i >= l; i-- {
        for j > mid && a[i] <= a[j] + diff {
            j--
        }
        ans += r - j
    }

    tmp := make([]int, r - l + 1)
    i, j, k := l, mid + 1, 0
    for i <= mid && j <= r {
        if a[i] <= a[j] {
            tmp[k] = a[i]
            i++
        } else {
            tmp[k] = a[j]
            j++
        }
        k++
    }

    for i <= mid {
        tmp[k] = a[i]
        k++
        i++
    }
    for j <= r {
        tmp[k] = a[j]
        k++
        j++
    }

    copy(a[l:r+1], tmp)
    return 
}