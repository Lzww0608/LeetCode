func minUnlockedIndices(nums []int, locked []int) (ans int) {
    if !check(nums) {
        return -1
    }
    n := len(nums)

    one, two := -1, n
    for i := 0; i < n; i++ {
        if nums[i] == 2 {
            two = i
            break
        }
    }
    for i := n - 1; i >= 0; i-- {
        if nums[i] == 1 {
            one = i 
            break
        }
    }
    for i := two; i < one; i++ {
        if locked[i] == 1 {
            ans++
        }
    }

    two, three := -1, n
    for i := n - 1; i >= 0; i-- {
        if nums[i] == 2 {
            two = i 
            break
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] == 3 {
            three = i
            break
        }
    }
    for i := three; i < two; i++ {
        if locked[i] == 1 {
            ans++
        }
    }

    return
}

func check(a []int) bool {
    n := len(a)
    one, three := -1, n

    for i := n - 1; i >= 0; i-- {
        if a[i] == 1 {
            one = i 
            break
        }
    }

    for i := 0; i < n; i++ {
        if a[i] == 3 {
            three = i
            break
        }
    }

    return three > one
}
