impl Solution {
    pub fn minimize_xor(num1: i32, num2: i32) -> i32 {
        let mut cnt1 = num1.count_ones();
        let mut cnt2 = num2.count_ones();

        let mut ans = num1;
        while cnt1 < cnt2 {
            cnt1 += 1;
            ans |= ans + 1;
        }
        while cnt1 > cnt2 {
            cnt1 -= 1;
            ans &= ans - 1
        }

        ans
    }
}