
import "fmt"
func subarrayLCM(nums []int, k int) (ans int) {
    n := len(nums)
    for i := 0; i < n; i++ {
        l := 1
        for j := i; j < n; j++ {
            l = l * nums[j] / gcd(l, nums[j])
            if l == k {
                //fmt.Println(i, j)
                ans++
            } else if k % l != 0 {
                break
            }
        }
    }

    return 
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}