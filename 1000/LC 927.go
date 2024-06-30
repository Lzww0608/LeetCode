func threeEqualParts(arr []int) []int {
    n, cnt := len(arr), 0
    for _, x := range arr {
        cnt += x
    }
    if cnt % 3 != 0 {
        return []int{-1, -1}
    } else if cnt == 0 {
        return []int{0, n - 1}
    }
    
    cnt /= 3

    find := func(x int) int {
        sum := 0
        for i, t := range arr {
            sum += t
            if sum == x {
                return i
            }
        }
        return 0
    }

    i, j, k := find(1), find(cnt + 1), find(cnt * 2 + 1)
    for k < n && arr[i] == arr[j] && arr[j] == arr[k] {
        i++
        j++
        k++
    }
    if k == n {
        return []int{i - 1, j}
    }

    return []int{-1, -1}
}
