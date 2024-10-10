use std::collections::HashMap;

impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i64 {
        let mut ans: i64 = 0;
        let mut cnt1 = HashMap::new();
        let mut mx = 0;

        for &x in &nums1 {
            if x % k == 0 {
                mx = mx.max(x / k);
                *cnt1.entry(x / k).or_insert(0) += 1;
            }
        }

        if mx == 0 {
            return 0;
        }

        let mut cnt2 = HashMap::new();

        for &x in &nums2 {
            *cnt2.entry(x).or_insert(0) += 1;
        }

        for (&key, &value) in &cnt2 {
            let mut sum = 0;
            let mut x = key;
            while x <= mx {
                sum += cnt1.get(&x).unwrap_or(&0);
                x += key;
            }

            ans += sum as i64 * value as i64;
        }

        ans
    }
}