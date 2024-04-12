func duplicateZeros(arr []int)  {
    n := len(arr)
    i, j := 0, 0
    for j < n {
        if arr[i] == 0 {
            j++
        } 
        i++
        j++
    }
    i--
    j--
    for i >= 0 {
        if j < n {
            arr[j] = arr[i]
        }
        if arr[i] == 0 {
            j--
            arr[j] = 0
        }
        i--
        j--
    }
}