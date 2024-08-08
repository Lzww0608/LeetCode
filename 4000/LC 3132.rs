impl Solution {
    pub fn minimum_added_integer(mut a: Vec<i32>, mut b: Vec<i32>) -> i32 {
        a.sort_unstable();
        b.sort_unstable();

        for i in (1..3).rev() {
            let x = b[0] - a[i];
            let mut j: usize = 0;
            for &v in &a[i..] {
                if b[j] == v + x && {j += 1; j == b.len()} {
                    return x;
                }
            }
        }

        b[0] - a[0]
    }
}