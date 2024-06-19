func countBeautifulPairs(nums []int) (ans int) {
    cnt := [10]int{}
    for _, x := range nums {
        for y := range cnt {
            if y > 0 && gcd(y, x % 10) == 1 {
                ans += cnt[y]
            }
        }
        for x >= 10 {
            x /= 10
        }
        cnt[x]++
    }

    return
}


func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }

    return x
}