impl Solution {
    pub fn semi_ordered_permutation(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let (mut l, mut r) = (0, 0);
        for i in 0..n {
            if nums[i] == 1 {
                l = i;
            } else if nums[i] == n as i32 {
                r = i;
            }
        }

        (l + n - r - 1 - (l > r) as usize) as _
    }
}