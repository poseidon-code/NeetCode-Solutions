// 42: Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int trap(vector<int>& height) {
        int l = 0;
        int r = height.size() - 1;

        int maxLeft = height[l];
        int maxRight = height[r];

        int result = 0;

        while (l<r) {
            if (maxLeft<=maxRight) {
                l++;
                maxLeft = max(maxLeft, height[l]);
                result += maxLeft - height[l];
            } else {
                r--;
                maxRight = max(maxRight, height[r]);
                result += maxRight - height[r];
            }
        }
        
        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> height = {4,2,0,3,2,5};
 
    // OUTPUT
    auto result = o.trap(height);
    cout<<result<<endl;

    return 0;
}