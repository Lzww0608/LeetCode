impl Solution {
    pub fn min_domino_rotations(tops: Vec<i32>, bottoms: Vec<i32>) -> i32 {
        let n = tops.len() as i32;

        let t = Self::solve(tops[0], &tops, &bottoms);
        let b = Self::solve(bottoms[0], &tops, &bottoms);

        if t == n && b == n {
            return -1;
        }

        std::cmp::min(t, b)
    }

    pub fn solve(target: i32, tops: &[i32], bottoms: &[i32]) -> i32 {
        let n = tops.len();
        let mut t = 0;
        let mut b = 0;

        for i in 0..n {
            if tops[i] != target && bottoms[i] != target {
                return n as i32;
            }

            if tops[i] != target {
                t += 1;
            }

            if bottoms[i] != target {
                b += 1;
            }
        }

        std::cmp::min(t, b)
    }
}