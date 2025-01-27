impl Solution {
    pub fn subset_xor_sum(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let n = nums.len() as i32;
        for &x in &nums {
            ans |= x;
        }
        ans << (n - 1)
    }
}