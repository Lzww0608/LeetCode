/**
 * Definition for an infinite stream.
 * type InfiniteStream interface {
 *     Next() int
 * }
 */

func kmp(a []int) []int {
    n := len(a)
    pi := make([]int, n)
    cnt := 0
    for i := 1; i < n; i++ {
        for cnt > 0 && a[cnt] != a[i] {
            cnt = pi[cnt - 1]
        }

        if a[cnt] == a[i] {
            cnt++
        }
        pi[i] = cnt 
    }

    return pi
}

func findPattern(stream InfiniteStream, pattern []int) int {
    n := len(pattern)
    pi := kmp(pattern)

    j := 0
    for i := 0; ; i++ {
        x := stream.Next()
        for j > 0 && pattern[j] != x {
            j = pi[j - 1]
        }
        if x == pattern[j] {
            j++
            if j == n {
                return i - j + 1
            }
        } 
    }

    return -1
}