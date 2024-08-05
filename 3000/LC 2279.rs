impl Solution {
    pub fn maximum_bags(capacity: Vec<i32>, rocks: Vec<i32>, additional_rocks: i32) -> i32 {
        let n = capacity.len();
        let mut additional_rocks = additional_rocks;
        let mut pos: Vec<usize> = (0..n).collect();

        pos.sort_by_key(|&i| capacity[i] - rocks[i]);

        let mut ans = 0;
        for &i in &pos {
            let t = capacity[i] - rocks[i];
            if t <= additional_rocks {
                additional_rocks -= t;
                ans += 1;
            } else {
                break;
            }
        }

        ans
    }
}