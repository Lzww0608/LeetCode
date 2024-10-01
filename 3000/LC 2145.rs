impl Solution {
    pub fn number_of_arrays(differences: Vec<i32>, lower: i32, upper: i32) -> i32 {
        let d = upper - lower;
        let (mut mx, mut mn) = (0, 0);
        let mut cur = 0;
        for &x in &differences {
            cur += x;
            mx = mx.max(cur);
            mn = mn.min(cur);
            if mx - mn > d {
                return 0;
            }
        }

        d - (mx - mn) + 1
    }
}