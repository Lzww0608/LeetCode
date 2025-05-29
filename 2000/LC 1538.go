/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 * // Compares 4 different elements in the array
 * // return 4 if the values of the 4 elements are the same (0 or 1).
 * // return 2 if three elements have a value equal to 0 and one element has value equal to 1 or vice versa.
 * // return 0 : if two element have a value equal to 0 and two elements have a value equal to 1.
 * func (this *ArrayReader) query(a, b, c, d int) int {}
 * 
 * // Returns the length of the array
 * func (this *ArrayReader) length() int {}
 */

func guessMajority(c *ArrayReader) int {
    n := c.length()
    a := make([]int, n)
    a[0] = 0
    cnt := [2]int{1, 0}

    solve := func(x, y, i, j int) {
        if x == y {
            a[j] = a[i]
        } else {
            a[j] = a[i] ^ 1
        }
        cnt[a[j]]++
    }

    solve(c.query(0, 2, 3, 4), c.query(1, 2, 3, 4), 0, 1)
    solve(c.query(0, 1, 3, 4), c.query(1, 2, 3, 4), 0, 2)
    solve(c.query(0, 1, 2, 4), c.query(1, 2, 3, 4), 0, 3)

    for i := 4; i < n; i++ {
        solve(c.query(0, 1, 2, 3), c.query(1, 2, 3, i), 0, i)
    }
    mx := 0
    if cnt[1] > cnt[0] {
        mx = 1
    } else if cnt[1] == cnt[0] {
        return -1
    }

    i := 0
    for i < n && a[i] != mx {
        i++
    }

    return i
}