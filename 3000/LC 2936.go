/**
 * Definition for BigArray.
 * type BigArray interface {
 *     At(int64) int
 *     Size() int64
 * }
 */
func countBlocks(nums BigArray) (ans int) {
    n := nums.Size()
    for i := int64(0); i < n; ans++ {
        x := nums.At(i)
        l, r := i, n
        for l < r {
            mid := l + ((r - l) >> 1)
            if nums.At(mid) != x {
                r = mid
            } else {
                l = mid + 1
            }
        }
        i = r
    }

    return
}