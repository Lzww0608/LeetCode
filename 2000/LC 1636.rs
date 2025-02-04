use std::collections::HashMap;
impl Solution {
    pub fn frequency_sort(mut nums: Vec<i32>) -> Vec<i32> {
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        for &x in nums.iter() {
            *cnt.entry(x).or_insert(0) += 1;
        }

        nums.sort_by(|a, b| {
            let cnt_a = cnt.get(a).unwrap();
            let cnt_b = cnt.get(b).unwrap();
            if cnt_a == cnt_b {
                b.cmp(a)
            } else {
                cnt_a.cmp(cnt_b)
            }
        });

        nums
    }
}