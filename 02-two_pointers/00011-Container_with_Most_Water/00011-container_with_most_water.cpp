// 11: Container with Most Water
// https://leetcode.com/problems/container-with-most-water/


#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    // SOLUTION
    int maxArea(vector<int>& height) {
        int l = 0;
        int r = height.size() - 1;
        int result = 0;

        while (l<r) {
            result = max(result, (r-l)*min(height[l], height[r]));
            if (height[l]<height[r]) l++;
            else r--;
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> height = {1,8,6,2,5,4,8,3,7};

    // OUTPUT
    auto result = o.maxArea(height);
    cout<<result<<endl;

    return 0;
}