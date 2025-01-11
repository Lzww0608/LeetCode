type SparseVector struct {
    a [][]int
}

func Constructor(nums []int) SparseVector {
    a := [][]int{}
    for i, x := range nums {
        if x != 0 {
            a = append(a, []int{i, x})
        }
    }
    return SparseVector{a}
}

// Return the dotProduct of two sparse vectors
func (c *SparseVector) dotProduct(vec SparseVector) (ans int) {
    m, n := len(c.a), len(vec.a)
    i, j := 0, 0
    for i < m && j < n {
        if c.a[i][0] == vec.a[j][0] {
            ans += c.a[i][1] * vec.a[j][1]
            i++
            j++
        } else if c.a[i][0] < vec.a[j][0] {
            i++
        } else {
            j++
        }
    }

    return
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */