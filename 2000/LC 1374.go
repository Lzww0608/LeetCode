把 k 个无区别的小球放到 n 个有区别的盒子中，允许盒子为空，一个盒子
import "strings"
func generateTheString(n int) string {
    if n & 1 == 1 {
        return strings.Repeat("a", n)
    }
    
    return strings.Repeat("a", n - 1) + "b"
}也可以放多个小球，有多少种不同的放法？