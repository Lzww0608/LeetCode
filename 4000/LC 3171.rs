use std::cmp::min;

impl Solution {
    pub fn minimum_difference(mut nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let mut ans = i32::MAX;
        for i in 0..n {
            ans = min(ans, (nums[i]-k).abs());
            let mut j = i - 1;
            while j < nums.len() && (nums[j] | nums[i]) != nums[j] {
                nums[j] |= nums[i];
                ans = min(ans, (nums[j] - k).abs());
                j -= 1;
            }
        }

        ans
    }
}