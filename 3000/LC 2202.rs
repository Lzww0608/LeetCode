use::std::cmp::{min, max};
impl Solution {
    pub fn maximum_top(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        if n == 1 && k & 1 == 1 {
            return -1;
        }

        let mut ans = 0;
        for i in 0..min(n, k as usize + 1) {
            if i != k as usize - 1 {
                ans = ans.max(nums[i]);
            }
        }

        ans
    }
}