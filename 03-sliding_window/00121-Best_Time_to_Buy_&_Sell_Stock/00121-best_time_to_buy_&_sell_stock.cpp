// 121. Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int maxProfit(vector<int>& prices) {
        int minValue = prices[0];
        int result = 0;

        for (int i=1; i<prices.size(); i++) {
            minValue = min(minValue, prices[i]);
            result = max(result, prices[i] - minValue);
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> prices = {7,1,5,3,6,4};

    // OUTPUT
    auto result = o.maxProfit(prices);
    cout<<result<<endl;

    return 0;
}