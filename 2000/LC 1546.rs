use std::collections::HashMap;

impl Solution {
    pub fn max_non_overlapping(nums: Vec<i32>, target: i32) -> i32 {
        let n = nums.len();
        let mut pre = 0;
        let mut m: HashMap<i32, usize> = HashMap::new();
        let mut f = vec![0; n + 1];
        m.insert(0, 0);

        for i in 0..n {
            pre += nums[i];
            if let Some(&v) = m.get(&(pre - target)) {
                f[i+1] = f[i].max(f[v] + 1);
            } else {
                f[i+1] = f[i];
            }
            m.insert(pre, i + 1);
        }

        f[n]
    }
}