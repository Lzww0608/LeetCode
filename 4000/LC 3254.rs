use std::cmp::{min, max};

impl Solution {
    pub fn results_array(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let n = nums.len();
        let k = k as usize;
        let mut ans = vec![0; n - k + 1];
        ans[0] = nums[0];

        for i in 1..n {
            if k > 1 && nums[i] != nums[i-1] + 1 {
                for j in max(0, i as i32 - k as i32 + 1) as usize..min(n - k + 1, i) {
                    ans[j] = -1;
                }
            } else if i >= k - 1 && ans[i+1-k] != -1 && nums[i] - nums[i+1-k] == k as i32 - 1 {
                ans[i+1-k] = nums[i];
            } 
            //println!("{:?}", ans);
        }

        ans
    }
}