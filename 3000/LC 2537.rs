use std::collections::HashMap;

impl Solution {
    pub fn count_good(nums: Vec<i32>, k: i32) -> i64 {
        let mut ans = 0;
        let n = nums.len();
        let mut cnt = HashMap::new();
        let mut r = 0;
        let mut sum = 0;

        for l in 0..n {
            while r < n && sum < k {
                let x = nums[r];
                *cnt.entry(x).or_insert(0) += 1;
                sum += cnt[&x] - 1;
                r += 1;
            }

            if sum >= k {
                ans += (n - r + 1) as i64;
                let x = nums[l];
                *cnt.entry(x).or_insert(0) -= 1;
                sum -= cnt[&x];
            } else {
                break
            }
        }

        ans
    }
}