const MAX int = 100_005
var f [MAX]int
func init() {
    for i := range f {
        sum, cnt := 0, 0
        for j := 1; j * j <= i; j++ {
            if i % j == 0 {
                if j * j == i || cnt >= 4 {
                    cnt = math.MaxInt32
                    break
                }
                cnt += 2;
                sum += j + i / j
            }
        }
        if cnt == 4 {
            f[i] = sum
        }
    }
}



func sumFourDivisors(nums []int) (ans int) {
    for _, x := range nums {
        ans += f[x]
    }

    return
}