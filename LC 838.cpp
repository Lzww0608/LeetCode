class Solution {
public:
    string pushDominoes(string dominoes) {
        int n = dominoes.length();
        int l = 0, r = 0;
        while (r < n && dominoes[r] == '.') r++;
        if (r < n && dominoes[r] == 'L') {
            while (l < r) {
                dominoes[l] = 'L';
                l++;
            }
        } else {
            l = r;
        }
        r++; 
        while (r < n) {
            while (r < n && dominoes[r] == '.') r++;
            if (r >= n) break;
            if (dominoes[r] == 'L' && dominoes[l] == 'R') {
                for (int i = l + 1, j = r - 1; i < j; i++, j--) {
                    dominoes[i] = 'R';
                    dominoes[j] = 'L';
                }
            } else if (dominoes[r] == 'R' && dominoes[l] == 'R') {
                for (int i = l + 1, j = r - 1; i <= j; i++, j--) {
                    dominoes[i] = 'R';
                    dominoes[j] = 'R';
                }
            } else if (dominoes[r] == 'L' && dominoes[l] == 'L') {
                for (int i = l + 1, j = r - 1; i <= j; i++, j--) {
                    dominoes[i] = 'L';
                    dominoes[j] = 'L';
                }
            } 
            l = r;
            r++;
        }
        if (l < n && dominoes[l] == 'R') {
            while (l < n) {
                dominoes[l] = 'R';
                l++;
            }
        }
        return dominoes;
    }
};