func inventoryManagement(stock []int) int {
    major, cnt := 0, 0
    for _, x := range stock {
        if major != x {
            cnt--
            if cnt <= 0 {
                major, cnt = x, 1
            }
        } else {
            cnt++
        }
    }

    return major
}
