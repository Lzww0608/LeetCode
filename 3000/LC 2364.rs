use std::collections::HashMap;

impl Solution {
    pub fn count_bad_pairs(nums: Vec<i32>) -> i64 {
        let mut ans = 0;
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        for (i, &x) in nums.iter().enumerate() {
            let t = x - i as i32;
            ans += i as i64 - *cnt.get(&t).unwrap_or(&0) as i64;
            *cnt.entry(t).or_insert(0) += 1;
        }        

        ans
    }
}