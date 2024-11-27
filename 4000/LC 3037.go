/**
 * Definition for an infinite stream.
 * type InfiniteStream interface {
 *     Next() int
 * }
 */
func findPattern(stream InfiniteStream, pattern []int) int {
    n := len(pattern)
    cnt := 0
    pi := make([]int, n)
    for i := 1; i < n; i++ {
        for cnt > 0 && pattern[cnt] != pattern[i] {
            cnt = pi[cnt-1]
        }
        if pattern[cnt] == pattern[i] {
            cnt++
        }
        pi[i] = cnt
    } 

    cnt = 0
    for i := 0; ; i++ {
        c := stream.Next()
        for cnt > 0 && pattern[cnt] != c {
            cnt = pi[cnt-1]
        }
        if pattern[cnt] == c {
            cnt++
        }
        if cnt == n {
            return i - n + 1
        }
    }

    return -1
}