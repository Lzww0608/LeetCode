use std::collections::HashMap;
use std::cmp::max;
impl Solution {
    pub fn maximum_length(nums: Vec<i32>, k: i32) -> i32 {
        let k = k as usize;
        let mut mx = vec![0; k + 2];
        let mut fs: HashMap<i32, Vec<i32>> = HashMap::new();

        for &x in &nums {
            let f = fs.entry(x).or_insert(vec![0; k + 1]);

            for x in (0..=k).rev() {
                f[x] = max(f[x], mx[x]) + 1;
                mx[x+1] = max(mx[x+1], f[x]);
            }
        }

        mx[k+1]
    }
}