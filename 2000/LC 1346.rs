use std::collections::HashMap;
impl Solution {
    pub fn check_if_exist(arr: Vec<i32>) -> bool {
        let mut cnt: HashMap<i32, bool> = HashMap::new();
        for &x in &arr {
            if cnt.contains_key(&(x * 2)) {
                return true;
            } else if x & 1 == 0 && cnt.contains_key(&(x / 2)) {
                return true;
            }
            cnt.insert(x, true);
        }

        false
    }
}