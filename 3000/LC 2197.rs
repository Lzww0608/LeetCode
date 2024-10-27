impl Solution {
    pub fn replace_non_coprimes(nums: Vec<i32>) -> Vec<i32> {
        let mut st: Vec<i32> = Vec::new();

        fn gcd(mut x: i64, mut y: i64) -> i64 {
            while y != 0 {
                let tmp = y;
                y = x % y;
                x = tmp;
            }
            x
        }

        for &x in &nums {
            st.push(x);

            while st.len() > 1 {
                let z = st[st.len()-1] as i64;
                let y = st[st.len()-2] as i64;
                let t = gcd(z, y);
                if t > 1 {
                    st.pop();
                    st.pop();
                    st.push((z * y / t) as i32);
                } else {
                    break
                }
            }
        }

        st
    }
}