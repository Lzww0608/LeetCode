func threeConsecutiveOdds(arr []int) bool {
    n := len(arr)
    for i := 2; i < n; i++ {
        if arr[i] & 1 == 1 {
            if arr[i-1] & 1 == 1 {
                if arr[i-2] & 1 == 1 {
                    return true
                } 
            } else {
                i++
            }
        } else {
            i += 2
        }
    }

    return false
}