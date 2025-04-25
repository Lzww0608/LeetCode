func transformArray(arr []int) []int {
    for true {
        f := true 
        tmp := make([]int, len(arr))
        copy(tmp, arr)
        for i := 1; i < len(arr) - 1; i++ {
            if tmp[i] > tmp[i-1] && tmp[i] > tmp[i+1] {
                arr[i]--
                f = false
            } else if tmp[i] < tmp[i-1] && tmp[i] < tmp[i+1] {
                arr[i]++
                f = false
            }
        }
        if f {
            break
        }
    }

    return arr
}