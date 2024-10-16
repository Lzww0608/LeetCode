use std::collections::HashMap;

impl Solution {
    pub fn num_of_pairs(nums: Vec<String>, target: String) -> i32 {
        let mut cnt = HashMap::new();
        let mut ans = 0;
        for num in &nums {
            *cnt.entry(num).or_insert(0) += 1;
        }

        let target = target.as_bytes();
        for i in 1..target.len() {
            let l = String::from_utf8(target[..i].to_vec()).unwrap();
            let r = String::from_utf8(target[i..].to_vec()).unwrap();

            if l != r {
                if let Some(&cnt_l) = cnt.get(&l) {
                    if let Some(&cnt_r) = cnt.get(&r) {
                        ans += cnt_l * cnt_r;
                    }
                }
            } else {
                if let Some(&cnt_l) = cnt.get(&l) {
                    ans += cnt_l * (cnt_l - 1);
                }
            }
        }

        ans
    }
}