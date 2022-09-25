// 875: Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/


#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;


class Solution {
public:
    // SOLUTION
    int minEatingSpeed(vector<int>& piles, int h) {
        int l = 1, r = *max_element(piles.begin(), piles.end());
        
        while (l < r) {
            int total = 0;
            int m = (l + r) / 2;

            for (int p : piles)
                total += (p + m - 1) / m;
            if (total > h)
                l = m + 1;
            else
                r = m;
        }

        return l;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> piles = {3,6,7,11};
    int h = 8;

    // OUTPUT
    auto result = o.minEatingSpeed(piles, h);
    cout<<result<<endl;

    return 0;
}