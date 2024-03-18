type NumArray struct {
    pre []int
}


func Constructor(nums []int) NumArray {
    pre := make([]int, len(nums) + 1)

    for i, x := range nums {
        pre[i + 1] = pre[i] + x
    }

    return NumArray{pre}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.pre[right + 1] - this.pre[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */