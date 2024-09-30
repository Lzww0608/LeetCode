use std::collections::HashMap;

impl Solution {
    pub fn find_lonely(nums: Vec<i32>) -> Vec<i32> {
        let mut cnt = HashMap::new();
        let mut ans = Vec::new();
        
        for &x in &nums {
            *cnt.entry(x).or_insert(0) += 1;
        }

        for (k, &v) in &cnt {
            if v == 1 && !cnt.contains_key(&(k-1)) && !cnt.contains_key(&(k+1)) {
                ans.push(*k);
            } 
        }

        ans
    }
}