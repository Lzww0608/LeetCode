impl Solution {
    pub fn count_collisions(s: String) -> i32 {
        let s = s.trim_start_matches('L').trim_end_matches('R');
        (s.len() - s.matches('S').count()) as i32
    }
}