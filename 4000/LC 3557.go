const N int = 26
func maxSubstrings(word string) (ans int) {
    pre := [N]int{}

    for i, c := range word {
        x := c - 'a'
        if pre[x] != 0 && i + 1 - pre[x] >= 3 {
            ans++
            pre = [N]int{}
        } else if pre[x] == 0 {
            pre[x] = i + 1
        }
    }

    return
}