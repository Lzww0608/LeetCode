//Go
func wiggleSort(nums []int)  {
    n := len(nums)
    arr := append([]int{}, nums...)
    sort.Ints(arr)
    i, j := (n + 1) / 2 - 1, n - 1
    for k := 0; k < n; k += 2 {
        nums[k] = arr[i]
        if k + 1 < n {
            nums[k+1] = arr[j]
        }
        i, j = i-1, j-1
    }
    return
}
//C++
