func findMagicIndex(nums []int) int {
    n := len(nums)

    var merge func(int, int) int 
    merge = func(l, r int) int {
        if l > r {
            return -1
        } 

        mid := l + ((r - l) >> 1)
        ans := merge(l, mid - 1)
        if ans != -1 {
            return ans 
        } else if nums[mid] == mid {
            return mid
        }

        return merge(mid + 1, r)
    }

    return merge(0, n - 1)
}