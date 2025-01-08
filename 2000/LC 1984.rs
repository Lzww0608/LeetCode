impl Solution {
    pub fn minimum_difference(mut nums: Vec<i32>, k: i32) -> i32 {
        nums.sort_unstable();
        let n = nums.len();
        let mut ans = nums[n-1];
        let k = k as usize;
        for i in 0..(n-k+1) {
            ans = ans.min(nums[i+k-1] - nums[i]);
        }

        ans
    }
}