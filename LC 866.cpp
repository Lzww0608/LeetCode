class Solution {
public:
    bool isPrime(int n) {
        if (n < 2) return false;
        for (int i = 2; i <= sqrt(n); i++) {
            if (n % i == 0) return false;
        }
        return true;
    }

    int reverse(int n) {
        int reversed = 0;
        while (n > 0) {
            reversed = 10 * reversed + (n % 10);
            n /= 10;
        }
        return reversed;
    }

    bool isPalindrome(int n) {
        return n == reverse(n);
    }



    int primePalindrome(int n) {
        while (true) {
            if (isPalindrome(n) && isPrime(n)) return n;
            n++;
            if (n > 11 && to_string(n).size() % 2 == 0) {
                n = pow(10, to_string(n).size());
            }
        }
    }

};