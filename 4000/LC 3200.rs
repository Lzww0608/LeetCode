impl Solution {
    pub fn max_height_of_triangle(red: i32, blue: i32) -> i32 {
        let (mut a, mut b) = (0, 0);
        let mut i = 1;
        loop {
            if (i & 1) == 0 {
                a += i;
            } else {
                b += i;
            }

            if (a > red || b > blue) && (a > blue || b > red) {
                return i - 1;
            }

            i += 1;
        }

        0
    }
}