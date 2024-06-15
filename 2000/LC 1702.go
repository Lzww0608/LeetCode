func maximumBinaryString(binary string) string {
    s := []byte(binary)
    n := len(binary)
    //cnt表示前置连续的1，无需更改
    cnt := 0
    for i := range s {
        if s[i] == '1' {
            cnt++
        } else {
            s[i] = '1'
            break
        }
    }

    if cnt == n {
        return binary
    }

    //sum表示cnt之后的1的个数，放在最后
    sum := 0
    for i := cnt + 1; i < n; i++ {
        if s[i] == '1' {
            sum++
        } else {
            s[i] = '1'
        }
    }
    //最多只有一个0
    s[n-sum-1] = '0'

    return string(s)
    
}
