use std::cmp::min;

impl Solution {
    pub fn losing_player(x: i32, y: i32) -> String {
        let t = min(x, y / 4);
        if t & 1 == 0 {
            return "Bob".to_string();
        }

        "Alice".to_string()
    }
}