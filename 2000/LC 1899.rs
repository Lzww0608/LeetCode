use std::cmp::max;

impl Solution {
    pub fn merge_triplets(triplets: Vec<Vec<i32>>, target: Vec<i32>) -> bool {
        let mut ans = vec![0; 3];
        for v in &triplets {
            if v[0] <= target[0] && v[1] <= target[1] && v[2] <= target[2] {
                for i in 0..3 {
                    ans[i] = max(ans[i], v[i]);
                }
            }
        }

        for i in 0..3 {
            if ans[i] != target[i] {
                return false;
            }
        }

        true
    }
}