use std::cmp::{max, min};
impl Solution {
    pub fn max_strength(nums: Vec<i32>) -> i64 {
        let mut mn = nums[0] as i64;
        let mut mx = mn;

        for i in 1..nums.len() {
            let x = nums[i] as i64;
            let tmp = mn; 
            mn = min(min(mn, x), min(mn * x, mx * x));
            mx = max(max(mx, x), max(tmp * x, mx * x));
        }

        mx
    }
}