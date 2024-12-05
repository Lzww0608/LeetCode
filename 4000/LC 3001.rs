impl Solution {
    pub fn min_moves_to_capture_the_queen(a: i32, b: i32, c: i32, d: i32, e: i32, f: i32) -> i32 {
        let in_between = |l: i32, m: i32, r: i32| l.min(r) < m && m < l.max(r);
        if a == e && (c != e || !in_between(b, d, f)) || 
            b == f && (d != f || !in_between(a, c, e)) ||
            c + d == e + f && (a + b != e + f || !in_between(c, a, e)) ||
            c - d == e - f && (a - b != c - d || !in_between(c, a, e)) {
                return 1;
            }
        2
    }
}