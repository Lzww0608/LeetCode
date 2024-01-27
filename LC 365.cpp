class Solution {
public:
    int gcd(int x, int y) {
        while (y) {
            int tmp = y;
            y = x % y;
            x = tmp;
        }
        return x;
    }
    bool canMeasureWater(int jug1Capacity, int jug2Capacity, int targetCapacity) {
        int x = gcd(jug1Capacity, jug2Capacity);
        if (targetCapacity % x == 0 && jug1Capacity + jug2Capacity >= targetCapacity) {
            return true;
        }
        return false;
    }
};
