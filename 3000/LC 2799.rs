use std::collections::{HashSet, HashMap};

impl Solution {
    pub fn count_complete_subarrays(nums: Vec<i32>) -> i32 {
        let mut m = HashSet::new();
        for &x in &nums {
            m.insert(x);
        }
        let cnt = m.len();

        let mut ans = 0;
        let mut l = 0;
        let mut vis = HashMap::new();
        for &x in &nums {
            *vis.entry(x).or_insert(0) += 1;
            
            while vis.len() == cnt {
                *vis.entry(nums[l]).or_insert(0) -= 1;
                if *vis.entry(nums[l]).or_insert(0) == 0 {
                    vis.remove(&nums[l]);
                }
                l += 1;
            }

            ans += l;
        }

        ans as _
    }
}