impl Solution {
    pub fn max_ascending_sum(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut cur = 0;
        let n = nums.len();

        for i in 0..n {
            if i == 0 || nums[i] > nums[i-1] {
                cur += nums[i];
            } else {
                cur = nums[i];
            }

            ans = ans.max(cur);
        }

        ans
    }
}