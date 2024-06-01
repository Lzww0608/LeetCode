func numDecodings(s string) int {
    const MOD int64 = 1000000007 // 修正模数为1000000007
    f := [3]int64{}
    var two, sum int64 = 1, 1
    for _, c := range s {
        if c == '*' {
            two = f[1] * 9 + f[2] * 6
            f[0], f[1], f[2] = sum * 7, sum, sum
        } else {
            x := int(c - '0')
            two = f[1]
            if x < 7 {
                two += f[2]
            }
            f = [3]int64{}
            if x == 1 || x == 2 {
                f[x] = sum
            } else if x != 0 {
                f[0] = sum
            }
        }
        sum = (f[0] + f[1] + f[2] + two) % MOD
    }

    return int(sum)
}
