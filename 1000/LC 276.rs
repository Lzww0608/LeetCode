impl Solution {
    pub fn num_ways(n: i32, k: i32) -> i32 {
        if n == 1 {
            return k;
        } else if n == 2 {
            return k * k;
        }

        let mut a = k;
        let mut b = k * k;
        for i in 2..n {
            let cur = (a + b) * (k - 1);
            a = b;
            b = cur;
        }

        b
    }
}