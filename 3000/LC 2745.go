func longestString(x int, y int, z int) int {
    // xx, xy, xz, yy, yx, yz, zz, zx, zy
    // xx, yy, xz, zy 
    // xy, yx, yz, zz, zx 
    ans := (z + min(x, y) * 2) * 2 
    if x != y {
        ans += 2
    }

    return ans
}