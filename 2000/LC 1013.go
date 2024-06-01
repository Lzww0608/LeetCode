func canThreePartsEqualSum(arr []int) bool {
    sum := 0
    for _, x := range arr {
        sum += x
    }

    if sum % 3 != 0 {
        return false
    }

    pre, cnt := 0, 0
    for _, x := range arr {
        pre += x 
        if pre == sum / 3 {
            pre = 0
            cnt++
        }
    }

    return cnt >= 3
}
