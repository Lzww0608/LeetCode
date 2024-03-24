func prisonAfterNDays(cells []int, n int) []int {

    bin := 0
    for i := range cells {
        bin = (bin << 1) | cells[i]
    }


    bin = ^(bin<<1 ^ bin>>1) & 0x7E
    start := bin

    memo := make([]int, 100)
    memo[0] = start


    for i := 1; i < n; i++ {
        bin = ^(bin<<1 ^ bin>>1) & 0x7E
        if bin == start { 
            bin = memo[(n-1)%i]
            break
        } else {
            memo[i] = bin
        }
    }

    for i := len(cells) - 1; i >= 0; i-- {
        cells[i] = bin & 1
        bin >>= 1
    }

    return cells
}

