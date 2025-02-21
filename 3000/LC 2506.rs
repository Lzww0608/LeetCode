use std::collections::HashMap;
impl Solution {
    pub fn similar_pairs(words: Vec<String>) -> i32 {
        let mut cnt: HashMap<i32, i32> = HashMap::new();
        let mut ans = 0;
        for s in &words {
            let mut t = s.as_bytes();
            let mut mask = 0;
            for i in 0..s.len() {
                mask |= 1 << (t[i] - b'a');
            }
            *cnt.entry(mask).or_insert(0) += 1;
            ans += *cnt.get(&mask).unwrap() - 1;
        }

        ans
    }
}