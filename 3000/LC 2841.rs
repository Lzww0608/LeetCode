use std::collections::HashMap;

impl Solution {
    pub fn max_sum(nums: Vec<i32>, m: i32, k: i32) -> i64 {
        let mut ans: i64 = 0;
        let mut cur: i64 = 0;
        let n = nums.len();

        let mut l = 0;
        let mut cnt = HashMap::new();
        for r in 0..n {
            cur += nums[r] as i64;
            *cnt.entry(nums[r]).or_insert(0) += 1;

            if r - l + 1 == k as usize {
                if cnt.len() >= m as _ {
                    ans = ans.max(cur);
                }

                cur -= nums[l] as i64;
                if let Some(count) = cnt.get_mut(&nums[l]) {
                    *count -= 1;
                    if *count == 0 {
                        cnt.remove(&nums[l]);
                    }
                }
                l += 1;
            }
        }

        ans
    }
}