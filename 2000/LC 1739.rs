impl Solution {
    pub fn minimum_boxes(n: i32) -> i32 {
        let (mut cur, mut i, mut j) = (0, 0, 0);
        while cur < n {
            j += 1;
            i += j;
            cur += i;
        }
        if cur == n {
            return i;
        }

        cur -= i;
        i -= j;
        j = 0;
        while cur < n {
            j += 1;
            cur += j;
        }

        i + j
    }
}