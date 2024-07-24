func minimumOperations(num string) int {
    n := len(num)
    ans := n - strings.Count(num, "0")

    flag := true
    for i, cnt := n - 1, 0; i >= 0; i-- {
        if num[i] == '0' {
            if flag {
                flag = false
            } else {
                ans = min(ans, cnt)
                break
            }
        } else if num[i] == '5' && !flag {
            ans = min(ans, cnt)
            break
        } else {
            cnt++
        }
    }

    flag = true
    for i, cnt := n - 1, 0; i >= 0; i-- {
        if num[i] == '5' && flag {
            flag = false
        } else if (num[i] == '7' || num[i] == '2') && !flag {
            ans = min(ans, cnt)
            break
        } else {
            cnt++
        }
    }


    return ans
}