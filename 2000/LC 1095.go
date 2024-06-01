/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, q *MountainArray) int {
    n := (*q).length()
    l, r, p := 1, n - 1, n / 2
    for l < r {
        mid := l + ((r - l) >> 1)
        a, b, c := (*q).get(mid-1), (*q).get(mid), (*q).get(mid+1)
        if b > a && b > c {
            p = mid
            break
        } else if b > a && c > b {
            l = mid + 1
        } else {
            r = mid
        }
    }
    //fmt.Println(p)

    find := func(l, r int, f bool) int {
        for l < r {
            mid := l + ((r - l) >> 1)
            t := (*q).get(mid)
            if t == target {
                return mid
            } else if t > target {
                if f {
                    r = mid
                } else {
                    l = mid + 1
                }
            } else if f {
                l = mid + 1
            } else {
                r = mid
            }
        }
        return -1
    }
    a := find(0, p + 1, true)
    if a != -1 {
        return a
    }
    return find(p + 1, n,false)
}
