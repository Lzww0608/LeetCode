/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    l, r := 0, 1_000_5
    for l < r {
        mid := l + ((r - l) >> 1)
        x := reader.get(mid)
        if x == target {
            return mid
        } else if x == (1 << 31) - 1 || x > target {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return -1
}