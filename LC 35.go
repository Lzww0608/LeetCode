func searchInsert(nums []int, target int) int {
    //return sort.SearchInts(nums, target)

    l, r := 0, len(nums)
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] > target {
            r = mid 
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            return mid 
        }
    }

    return l
}