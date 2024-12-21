
import "strconv"
func generateAbbreviations(word string) (ans []string) {
    n := len(word)

    var dfs func(int, string, int)
    dfs = func(i int, s string, a int) {
        if i == n {
            ans = append(ans, s) 
            return 
        }

        k := 1
        for j := i; j < n; j++ {
            if a != 1 {
                dfs(j + 1, s + word[i:j+1], 1)
            } 
            if a != 2 {
                t := strconv.Itoa(k)
                dfs(j + 1, s + t, 2)
                k++
            }
            
        }
    }

    dfs(0, "", 0)
    return
}