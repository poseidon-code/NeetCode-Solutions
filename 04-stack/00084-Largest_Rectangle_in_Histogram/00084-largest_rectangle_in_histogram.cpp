// 84: Largest Rectangle in Histogram
// https://leetcode.com/problems/largest-rectangle-in-histogram/


#include <iostream>
#include <vector>
#include <stack>

using namespace std;


class Solution {
public:
    // SOLUTION
    int largestRectangleArea(vector<int>& heights) {
        stack<int> stack;
        int l = heights.size();
        int result = 0;

        for (int i=0; i<=l; i++) {
            int h = i==l ? 0 : heights[i];

            if (stack.empty() || h>=heights[stack.top()])
                stack.push(i);
            else {
                int tp = stack.top();
                stack.pop();
                result = max(result, heights[tp]*(stack.empty() ? i : i-1-stack.top()));
                i--;
            }
        }

        return result;
    }
};


int main() {
    Solution o;

    // INPUT
    vector<int> heights = {2,1,5,6,2,3};

    // OUTPUT
    auto result = o.largestRectangleArea(heights);
    cout<<result<<endl;

    return 0;
}