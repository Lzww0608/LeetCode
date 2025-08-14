func canReorderDoubled(arr []int) bool {
    cnt := make(map[int]int)
    for _, x := range arr {
        cnt[x]++
    }
    sort.Ints(arr)
    n := len(arr)
    i := 0

    for ; i < n && arr[i] < 0; i++ {
        if cnt[arr[i]] == 0 {
            continue
        }
        
        if arr[i] & 1 == 1 || cnt[arr[i] / 2] < cnt[arr[i]] {
            return false;
        }
        cnt[arr[i] / 2] -= cnt[arr[i]];
        cnt[arr[i]] = 0
    }

    if i < n && arr[i] == 0 {
        if cnt[arr[i]] & 1 == 1 {
            return false
        }
        i++
    }

    for ; i < n; i++ {
        if cnt[arr[i]] == 0 {
            continue
        }

        if cnt[arr[i] * 2] < cnt[arr[i]] {
            return false
        }
        cnt[arr[i] * 2] -= cnt[arr[i]]
        cnt[arr[i]] = 0
    }

    return true
}