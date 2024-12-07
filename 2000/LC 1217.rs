impl Solution {
    pub fn min_cost_to_move_chips(position: Vec<i32>) -> i32 {
        let mut odd = 0;
        let mut even = 0;
        for &x in &position {
            if x & 1 == 0 {
                even += 1;
            } else {
                odd += 1;
            }
        }

        std::cmp::min(even, odd)
    }
}