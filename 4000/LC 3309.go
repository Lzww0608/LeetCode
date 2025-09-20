func maxGoodNumber(nums []int) int {
    a, b, c := nums[0], nums[1], nums[2]
    x, y, z := bits.Len(uint(a)), bits.Len(uint(b)), bits.Len(uint(c))

    ans := (a << (y + z)) + (b << z) + c 
    ans = max(ans, (a << (y + z)) + (c << y) + b)
    ans = max(ans, (b << (x + z)) + (c << x) + a)
    ans = max(ans, (b << (x + z)) + (a << z) + c)
    ans = max(ans, (c << (x + y)) + (b << x) + a)
    ans = max(ans, (c << (x + y)) + (a << y) + b)

    return ans
}