func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    m := map[int]int{}
    for _, v := range reservedSeats {
        if v[1] == 10 || v[1] == 1 {
            continue
        }
        m[v[0]] |= 1 << ((v[1] - 2) / 2)
    }

    ans := n << 1
    for _, x := range m {
        ans -= 2
        if (x & 0b0011) == 0 || (x & 0b0110) == 0 || (x & 0b1100) == 0 {
            ans++
        }
    }

    return ans
}