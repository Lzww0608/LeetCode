/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 1, n + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        f := guess(mid) 
        if f == 1 {
            l = mid + 1
        } else if f == -1 {
            r = mid 
        } else {
            return mid
        }
    }

    return l
}
