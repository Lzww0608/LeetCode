impl Solution {
    pub fn decompress_rl_elist(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut ans = Vec::new();
        for i in (0..n).step_by(2) {
            let freq = nums[i];
            let val = nums[i+1];
            for j in (0..freq) {
                ans.push(val);
            }
        }

        ans
    }
}