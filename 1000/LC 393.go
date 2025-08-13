func validUtf8(data []int) bool {
    cnt := 0
    for _, x := range data {
        if cnt == 0 {
            if (x >> 5) == 0b110 {
                cnt = 1
            } else if (x >> 4) == 0b1110 {
                cnt = 2
            } else if (x >> 3) == 0b11110 {
                cnt = 3
            } else if (x >> 7) == 1 {
                return false
            }
        } else {
            if (x >> 6) != 0b10 {
                return false
            }
            cnt--
        }
    }

    return cnt == 0
}