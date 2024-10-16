use std::collections::HashMap;

impl Solution {
    pub fn num_triplets(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let mut cnt1 = HashMap::new();
        let mut cnt2 = HashMap::new();
        let n = nums1.len();
        let m = nums2.len();
        let mut ans = 0;

        for i in 0..n {
            for j in (i + 1)..n {
                let square = nums1[i] as i64 * nums1[j] as i64;
                *cnt1.entry(square).or_insert(0) += 1;
            }
        }

        for i in 0..m {
            for j in (i + 1)..m {
                let square = nums2[i] as i64 * nums2[j] as i64;
                *cnt2.entry(square).or_insert(0) += 1;
            }
            let x = nums2[i] as i64 * nums2[i] as i64;
            if let Some(&count) = cnt1.get(&x) {
                ans += count;
            }
        }

        for i in 0..n {
            let square = nums1[i] as i64 * nums1[i] as i64;
            if let Some(&count) = cnt2.get(&square) {
                ans += count;
            }
        }

        ans
    }
}